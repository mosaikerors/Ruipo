package main

import (
	"database/sql"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)

func longLinkToShortLink(longLink string) string {
	//直接生成一个短链接,两种可能会冲突，一种就随机到一起去了，这个概率很小；第二种就是两个人同时执行到这了，随机种子一样
	rand.Seed(time.Now().Unix())
	shortLink := generateShortLink(6)
	db, _ := sql.Open("mysql", "root:1234567@tcp(3.92.33.121:3306)/link?charset=utf8")
	_, _ = db.Query("insert into link_table (shortLink, longLink) values (?, ?);", shortLink, longLink)
	_ = db.Close()
	return shortLink
}

func shortLinkToLongLink(shortLink string) string {
	//先去redis看有没有
	c, err := redis.Dial("tcp", "3.92.33.121:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return ""
	}
	defer c.Close()
	longLink, err := redis.String(c.Do("GET", shortLink))

	if longLink == "" {
		//redis中没有，去mysql拿
		db, _ := sql.Open("mysql", "root:1234567@tcp(3.92.33.121:3306)/link?charset=utf8")
		row:= db.QueryRow("select longLink from link_table where shortLink = ? limit 1;", shortLink)
		_ = row.Scan(&longLink)
		_ = db.Close()

		//插入到redis中
		_, err = c.Do("SET", shortLink, longLink)
		if err != nil {
			fmt.Println("redis set failed:", err)
		}
	} else {
		//redis中有，直接返回
		return longLink
	}

	return longLink
}
