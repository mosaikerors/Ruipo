FROM golang:latest as builder
WORKDIR /src/main
ADD . /src/main
RUN go get github.com/garyburd/redigo/redis && \
    go get github.com/go-sql-driver/mysql && \
    go get gopkg.in/yaml.v2
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest as prod
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /src/main/ .
CMD ["./app"]
