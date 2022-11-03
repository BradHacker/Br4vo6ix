package main

import (
	"encoding/base64"
	"net"
	"strconv"
	"strings"
)

// GetIP gets the IP of the local machine
func GetIP() string {
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

// DecodeProxyIps takes a base64 encoded, comma separated list of IP addresses and returns a slice of IPs
func DecodeProxyIps(encodedProxyIps string) ([]string, error) {
	proxyIpList, err := base64.StdEncoding.DecodeString(encodedProxyIps)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(string(proxyIpList), ","), nil
}

// DecodeProxyIps takes a base64 encoded, comma separated list of port numbers and returns a slice of port numbers
func DecodeProxyPorts(encodedProxyPorts string) ([]int, error) {
	proxyPortList, err := base64.StdEncoding.DecodeString(encodedProxyPorts)
	if err != nil {
		return []int{}, err
	}
	portStrings := strings.Split(string(proxyPortList), ",")
	ports := make([]int, len(portStrings))
	for i, p := range portStrings {
		if ports[i], err = strconv.Atoi(p); err != nil {
			ports[i] = 0
			continue
		}
	}
	return ports, nil
}
