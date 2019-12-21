package main

import (
	"fmt"
)

func insertToRedis() {
	for true {
		pair := <-insertRedisChannel
		fmt.Println(pair.linkIndex)
		_, _ = cluster.Do("SET", pair.linkIndex, pair.longLink)
	}
}
