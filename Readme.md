# mytholojam

## Usage

Docker repos is available at https://hub.docker.com/r/aopal/mytholojam

### Running the client


```bash
docker pull aopal/mytholojam:client-latest
server="http://host.docker.internal:8080" # configure
docker run -it aopal/mytholojam:client-latest $server
```

### Running the server

```bash
docker pull aopal/mytholojam:server-latest
port=8080 # configure
docker run -d -p $port:8080 aopal/mytholojam:server-latest
```


## Development

Requires golang https://golang.org/

### Running the server

Build & run server: `./s`

Build & run server docker container `script/docker-start-server`

Both scripts accept an optional argument to set the port of the server. Default port is `8080` otherwise

### Running the command line client

Build & run client:`./c`

Build & run client docker container `script/docker-start-client`

Both scripts accept and optional argument to set the URL of the server. Default server address is `http://localhost:8080` / `http://host.docker.internal:8080`
