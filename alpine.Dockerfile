FROM alpine:latest
RUN apk add --no-cache curl

ENTRYPOINT [ "sleep", "infinity" ]