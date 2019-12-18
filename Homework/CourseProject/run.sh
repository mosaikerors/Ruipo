#create database
cd database
docker build -t mysql_docker .
docker run -itd --name shortlink_mysql  -p 3306:3306 mysql_docker
docker run -p 6379:6379 -d redis:latest redis-server

#create service
cd ../main
docker build -t server_docker .
docker run -itd --name shortlink_server  -p 9090:9090 server_docker
