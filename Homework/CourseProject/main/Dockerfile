FROM golang:latest
WORKDIR $GOPATH/src/main
ADD . $GOPATH/src/main
RUN go get github.com/garyburd/redigo/redis && \
    go get github.com/go-sql-driver/mysql && \
    go get gopkg.in/yaml.v2
RUN go build .
EXPOSE 9090
ENTRYPOINT  ["./main"]