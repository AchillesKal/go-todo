FROM golang:1.8

RUN mkdir -p /go/src/go_todo
WORKDIR /go/src/go_todo

ADD . /go/src/go_todo

RUN go get -v