# Br4vo6ix

> **_DISCLAIMER: This tool is for educational, competition, and training purposes only. I am in no way responsible for any abuse of this tool_**

This is a golang C2 + Implant that communicates via [Protocol Buffers](https://developers.google.com/protocol-buffers) (aka. `protobuf`s).

> _Note: this tool is still somewhat in development_

## Standing up the C2 Server

### Pwnboard

This C2 server is designed to communicate with [Pwnboard](https://github.com/micahjmartin/pwnboard) for competition use. To tell Br4vo6ix where Pwnboard is being hosted, you must pass the url in the `PWN_URL` environment variable.

```shell
# I place this line in a .source file and just source it for ease-of-use, but
#   this just needs to be set in the terminal that runs the C2
export PWN_URL=http(s)://<url for pwnboard>/generic
```

### Frontend

The frontend needs to know where the GraphQL endpoint is located at, so you must create a `.env` file in the `frontend` folder with the following value:

```shell
# frontend/.env
REACT_APP_GRAPHQL_URL=http://<Br4vo6x host fqdn/ip>:8080/query
```

To compile the frontend (written in ReactJS), you can simply move into the `frontend` directory and run `npm run build`

```shell
$ cd frontend
$ npm run build
```

Typically the frontend is hosted via [Nginx](https://www.nginx.com/). This is the base template to use:

```
server {
    listen <PORT> default_server;
    listen [::]:<PORT> default_server;

    root <PATH TO BR4VO6IX>/ui/build;

    index index.html index.htm index.nginx-debian.html;

    server_name _;

    location / {
        try_files $uri /index.html;
    }
    location /api {
            proxy_pass http://localhost:<GRAPHQL PORT (default is 8080)>/api;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection $http_connection;
            proxy_http_version 1.1;
        }

}
```

## Compiling the implant

First, you will need an `.env` file to configure the compiled implants. This goes in the root directory.

```shell
# .env
REDTEAM_IPS=<IPS OF ALL PROXY SERVERS (COMMA SEPARATED; NO SPACES)>
PORTS=<IPS OF ALL PROXY PORTS (COMMA SEPARATED; NO SPACES)>
C2_IP=<IP OF C2 SERVER (BASE64 ENCODED)>
PROXY_IP=<IP OF C2 SERVER>
XOR_KEY=<SHARED XOR KEY (ANY STRING)>

WINDOWS_SRV_NAME=<NAME OF WINDOWS SERVICE>
WINDOWS_SRV_DESC=<DESCRIPTION OF WINDOWS SERVICE>
WINDOWS_OUT_FILE=<WINDOWS EXECUTABLE OUTPUT PATH>

LINUX_SRV_NAME=<NAME OF LINUX SERVICE>
LINUX_SRV_DESC=<DESCRIPTION OF LINUX SERVICE>
LINUX_OUT_FILE=<LINUX ELF BINARY OUTPUT PATH>

MAC_SRV_NAME=<NAME OF MACOS LAUNCHCTL SERVICE>
MAC_SRV_DESC=<DESCRIPTION OF MACOS LAUNCHCTL SERVICE>
MACOS_OUT_FILE=<MACOS MACH-O BINARY OUTPUT PATH>
```

Then, to compile the implant you should be able to run `make all`

## Proxy Servers

The C2 is designed to sit behind as many proxy servers as you want. We utilize `socat` for our proxy servers, so please install the `socat` package on all proxy server boxes.

The `make all` command will generate proxy scripts in the `out/scripts` directory. The scripts will be labeled `<IP ADDR>-proxy.sh` and `<IP ADDR>-unproxy.sh`. The IP Addresses match to each proxy server and should be run on the proxy servers in order to automatically configure `socat`.
