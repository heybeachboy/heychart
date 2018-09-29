FROM golang:latest
RUN go get github.com/gin-gonic/gin
MAINTAINER Razil "google@gmail.com"
WORKDIR $GOPATH/src/heychart
ADD . $GOPATH/src/heychart
RUN go get github.com/gin-gonic/gin
RUN go build .

EXPOSE 9090

ENTRYPOINT ["./heychart"]
