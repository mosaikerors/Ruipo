package main

import (
	"fmt"
	"strconv"

	"github.com/chasex/redis-go-cluster"
	_ "github.com/go-sql-driver/mysql"
)

func longLinkToShortLink(longLink string) string {
	var linkIndex int64 = 0

	tx, _ := db.Begin()
	result, _ := tx.Exec("insert into link (long_link) values (?);", longLink)
	linkIndex, _ = result.LastInsertId()
	_ = tx.Commit()

	return longIntToShortLink(linkIndex)
}

func shortLinkToLongLink(shortLink string) string {
	// convert short link to index
	linkIndex := shortLinkToLongInt(shortLink)

	// check if redis has cached it
	key := strconv.FormatInt(linkIndex, 10)
	fmt.Println("use strconv format")
	fmt.Println(key)
	longLink, rediserror := redis.String(cluster.Do("GET", key))
	if rediserror != nil {
		fmt.Println(rediserror.Error())
	}
	fmt.Println("in shortLinkToLongLink after read longLink in redis")
	fmt.Println(longLink)
	if longLink == "" {
		// select by link index
		row := db.QueryRow("select long_link from link where link_index = ? limit 1;", linkIndex)
		rowerr := row.Scan(&longLink)
		if rowerr != nil {
			fmt.Println(rowerr.Error())
		}
		//send to channel to tell -- insert to redis
		fmt.Println("find in mysql")
		fmt.Println(linkIndex)
		insertRedisChannel <- linkIndex2LongLinkPair{linkIndex: linkIndex, longLink: longLink}
	}
	return longLink
}
