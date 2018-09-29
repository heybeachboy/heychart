FROM golang:latest
MAINTAINER Razil "google@gmail.com"
WORKDIR $GOPATH/src/heychart
ADD . $GOPATH/src/heychart
RUN go build .

EXPOSE 9090

ENTRYPOINT ["./heychart"]
