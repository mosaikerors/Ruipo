package main

import (
	"fmt"
	"log"
	"net/http"
)

type linkIndex2LongLinkPair struct {
	linkIndex int64
	longLink string
}

var insertRedisChannel chan linkIndex2LongLinkPair
var mysqlConfig string
var redisConfig string

func main() {
	//read config.yml and parse
	parseConfig()
	fmt.Println(mysqlConfig)
	fmt.Println(redisConfig)
	insertRedisChannel = make(chan linkIndex2LongLinkPair, 100)
	go insertToRedis()
	http.HandleFunc("/short", handleShort)
	http.HandleFunc("/long", handleLong)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
