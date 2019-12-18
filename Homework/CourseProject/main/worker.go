package main

import "github.com/garyburd/redigo/redis"

func insertToRedis()  {
	c, _ := redis.Dial("tcp", redisConfig)
	for true {
		pair := <- insertRedisChannel
		_, _ = c.Do("SET", pair.linkIndex, pair.longLink)
	}
	c.Close()
}