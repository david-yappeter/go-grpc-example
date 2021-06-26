FROM golang:1.14-alpine as builder

WORKDIR /demo/

COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /demo/myapp .

FROM alpine:latest
COPY --from=builder /demo/myapp /demo/myapp
ENTRYPOINT [ "/demo/myapp" ]
