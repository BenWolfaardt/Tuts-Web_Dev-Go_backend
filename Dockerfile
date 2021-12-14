FROM golang:1.17.2-alpine
RUN mkdir /build
ADD go.mod go.sum hello.go /build/
WORKDIR /build
RUN go build