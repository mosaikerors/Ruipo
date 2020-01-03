# JMeter Notes

## 安装 Java

[centos 安装 java](<https://segmentfault.com/a/1190000015389941>)

*fg: 直接用 `yum install maven` 不仅可以装上 java8，还能直接装 maven，避免以后装 maven 的时候可能会有的版本冲突。*

## 配置 JMeter 环境

要进行[分布式压力测试](<https://jmeter.apache.org/usermanual/jmeter_distributed_testing_step_by_step.html>)的话，需要准备一个 master 机器和若干 slave 机器。

将 JMeter 安装包分发到 master 和 slave 节点，更改配置文件 `bin/jmeter.properties`:

```properties
# Set this if you don't want to use SSL for RMI, 一般我们测试情况就不开SSL
server.rmi.ssl.disable=true 
# 有关于你远端调用的IP or DNS
remote_hosts=node2,node3,node4
# 远端服务具体的port，最好指定，不然会是随机的端口，防火墙不好设置
server.rmi.localport=3030
# 有的博文里面说改成false，但是这样我的jmeter-server就启动不了，原因不明。
server.rmi.create=true
```

*[reference](<https://github.com/tx19980520/new-tech-stack/blob/master/Jmeter/Apache%20JMeter%20Distributed%20Testing.md>)*

然后在 slave 节点启动 jmeter-server，等待 master 调度:

```bash
./jmeter-server
Using local port: 3030
Created remote object: UnicastServerRef2 [liveRef: [endpoint:[10.0.0.24:3030](local),objID:[-300bb601:16d8b3a7c1c:-7fff, 12160871764474459]]]
```

## 启动测试

在 master 节点上启动测试：

```bash
./jmeter -n -t path/to/jmx/test.jmx -R node4,node3,node2 -l log.jtl -e -o dashboard
```

+ `-n` 表示用 CLI 启动测试
+ `-t` 表示 jmx 文件的位置（提前准备好）
+ `-R` 表示 slave 节点的 hostname，也可以用 `-r`，这样 slave 节点即为 `jmeter.properties` 中指定的那些
+ `-l` 表示 jtl 文件生成的位置，不能指定一个已存在的文件名
+ `-e` 表示生成 report dashboard
+ `-o` 表示 report dashboard 生成的位置，不能指定一个已存在且不为空的目录

##### Last-modified date: 2020.1.3, 4 p.m.