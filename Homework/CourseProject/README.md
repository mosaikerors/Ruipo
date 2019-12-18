## 部署

提供两种部署方式
- 一键部署
  ```shell
  ./run.sh
  ```
- 数据库、服务依次部署
  - 数据库
      ```shell
      cd database
      ```
      
      ```shell
      docker build -t mysql_docker .
      ```
      ```shell
      docker run -itd --name shortlink_mysql  -p 3306:3306 mysql_docker
      ```
      ```shell
      docker run -p 6379:6379 -d redis:latest redis-server
      ```
  - 服务
      ```shell
      cd main
      ```
      ```shell
      docker build -t server_docker .
      ```  
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