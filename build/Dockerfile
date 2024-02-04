FROM golang:1.21.6-alpine3.19

RUN apk update \
    && apk add protoc \
    && apk cache clean \
    && rm -rf /var/lib/apt/lists/*

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest