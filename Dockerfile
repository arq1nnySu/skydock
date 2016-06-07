FROM golang:1.6

# go get to download all the deps
RUN go get -u github.com/malyhov/skydock

ADD . /go/src/github.com/malyhov/skydock
ADD plugins/ /plugins

RUN cd /go/src/github.com/malyhov/skydock && go install . ./...

ENTRYPOINT ["/go/bin/skydock"]
