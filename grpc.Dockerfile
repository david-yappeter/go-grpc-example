FROM golang:1.16-alpine as builder
WORKDIR /demo/

COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /demo/myapp ./grpc/grpc_server.go

FROM alpine:latest
COPY --from=builder /demo/myapp /demo/myapp
EXPOSE 8080
ENTRYPOINT [ "/demo/myapp" ]
