package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type MyConfig struct {
	Mysql MysqlConfig
	Redis RedisConfig
}

type MysqlConfig struct {
	Ipport   string
	Username string
	Password string
}

type RedisConfig struct {
	Ipport string
}

func parseConfig() {
	data, _ := ioutil.ReadFile("./config/config.yml")
	t := MyConfig{}
	//turn yaml into struct
	_ = yaml.Unmarshal(data, &t)
	mysqlConfig = fmt.Sprintf("%s:%s@tcp(%s)/shortlink?charset=utf8",
		t.Mysql.Username, t.Mysql.Password, t.Mysql.Ipport)
	redisConfig = t.Redis.Ipport
}
