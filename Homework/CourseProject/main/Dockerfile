FROM golang:latest
WORKDIR $GOPATH/src/main
ADD . $GOPATH/src/main
RUN go get github.com/garyburd/redigo/redis
RUN go get github.com/go-sql-driver/mysql
RUN go build .
EXPOSE 9090
ENTRYPOINT  ["./main"]