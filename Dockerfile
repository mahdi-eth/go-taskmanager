FROM golang

ADD . /go/src/github.com/mahdi-eth/go-taskmanager

WORKDIR /go/src/github.com/mahdi-eth/go-taskmanager

RUN go install github.com/mahdi-eth/go-taskmanager

ENTRYPOINT /go/bin/taskmanager

EXPOSE 8080