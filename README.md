## go-GRPC Example

##### Convert .proto to go
docker-compose.yml
```
version: "3.5"

services:
    # Generate .proto to go
    protoc:
        image: "namely/protoc-all:latest"
        volumes: 
            - "./:/defs"
        command: -f ./rpc/testing/*.proto -l go -o .
```
```
$ docker-compose run protoc
```

##### Run Grpc Server
```
$ go run grpc/grpc_server.go
```

##### Run Main Code
```
$ go run server.go
```

open http://localhost:8080 (simple multiplication)

