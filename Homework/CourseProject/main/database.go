package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/garyburd/redigo/redis"
)

func longLinkToShortLink(longLink string) string {
	var linkIndex int64 = 0
	db, _ := sql.Open("mysql", mysqlConfig)

	tx, _ := db.Begin()
	_, _ = tx.Exec("insert into link (long_link) values (?);", longLink)

	row := tx.QueryRow("select count(*) from link;")
	err := row.Scan(&linkIndex)

	if err != nil{
		panic(err)
	}

	_ = tx.Commit()
	_ = db.Close()

	return longIntToShortLink(linkIndex)
}

func shortLinkToLongLink(shortLink string) string {
	// convert short link to index
	linkIndex := shortLinkToLongInt(shortLink)

	// check if redis has cached it
	c, _ := redis.Dial("tcp", redisConfig)
	defer c.Close()
	longLink, _ := redis.String(c.Do("GET", shortLink))

	if longLink == "" {
		// open mysql
		db, _ := sql.Open("mysql", mysqlConfig)
		// select by link index
		row:= db.QueryRow("select long_link from link where link_index = ? limit 1;", linkIndex)
		_ = row.Scan(&longLink)
		_ = db.Close()
		//send to channel to tell -- insert to redis
		insertRedisChannel <- linkIndex2LongLinkPair{linkIndex: linkIndex, longLink: longLink}
	}
	return longLink
}
