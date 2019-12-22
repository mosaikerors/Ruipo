package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/chasex/redis-go-cluster"
)

type linkIndex2LongLinkPair struct {
	linkIndex int64
	longLink  string
}

var insertRedisChannel chan linkIndex2LongLinkPair
var mysqlConfig string
var redisConfig string
var db *sql.DB
var cluster *redis.Cluster

func main() {
	//read config.yml and parse
	parseConfig()
	fmt.Println(mysqlConfig)
	fmt.Println(redisConfig)
	// init db pool
	db, _ = sql.Open("mysql", mysqlConfig)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	// init redis pool
	cluster, _ = redis.NewCluster(
		&redis.Options{
			StartNodes:   []string{redisConfig},
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})
	insertRedisChannel = make(chan linkIndex2LongLinkPair, 100)
	go insertToRedis()
	http.HandleFunc("/short", handleShort)
	http.HandleFunc("/long", handleLong)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
