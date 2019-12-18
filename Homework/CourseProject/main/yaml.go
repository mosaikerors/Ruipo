package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type MyConfig struct {
	Mysql MysqlConfig
	Redis RedisConfig
}

type MysqlConfig struct {
	Ipport    string
	Username string
	Password string
}

type RedisConfig struct {
	Ipport    string
}

func parseConfig()  {
	data, _ := ioutil.ReadFile("config.yml")
	t := MyConfig{}
	//把yaml形式的字符串解析成struct类型
	yaml.Unmarshal(data, &t)
	mysqlConfig = fmt.Sprintf("%s:%s@tcp(%s)/shortlink?charset=utf8",
		t.Mysql.Username, t.Mysql.Password, t.Mysql.Ipport)
	redisConfig = t.Redis.Ipport
}