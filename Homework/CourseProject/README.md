## 数据库

### 构建MYSQL数据库

```shell
cd database
```

```shell
docker build -t mysql_docker .
```

### 运行MYSQL数据库

```shell
docker run -itd --name shortlink_mysql  -p 3306:3306 mysql_docker
```

### 构建运行REDIS数据库

```shell
docker run -p 6379:6379 -d redis:latest redis-server
```

## 服务

注意事项

- 目前的go文件中，数据库的配置是写死的。在对服务进行打包之前，请到`database.go`文件中更改MYSQL数据库的IP和端口号；

### 构建服务镜像

```shell
cd main
```

```shell
docker build -t server_docker .
```

### 运行服务

```shell
docker run -itd --name shortlink_server  -p 9090:9090 server_docker
```

## 测试

注意事项

- 由于没有容错处理，请保证短链接变成长链接时，短链接一定存在；

### 长链接 -> 短链接

```shell
curl -X POST http://ip:9090/long?long_link=[your long link]
```

### 短链接 -> 长链接

```shell
curl -X POST http://ip:9090/short?short_link=[your short link]
```



