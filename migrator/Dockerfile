FROM alpine:latest

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

COPY ./goose_linux_x86_64 /bin/goose

RUN chmod +x /bin/goose

WORKDIR /migrations