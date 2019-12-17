package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func longLinkToShortLink(longLink string) string {
	var linkIndex int64 = 0
	db, _ := sql.Open("mysql",
		"root:123456@tcp(3.91.25.91:3306)/shortlink?charset=utf8")

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

	var longLink string
	// convert short link to index
	linkIndex := shortLinkToLongInt(shortLink)
	// open mysql
	db, _ := sql.Open("mysql",
		"root:123456@tcp(3.91.25.91:3306)/shortlink?charset=utf8")
	// select by link index
	row:= db.QueryRow("select long_link from link where link_index = ? limit 1;", linkIndex)
	_ = row.Scan(&longLink)
	_ = db.Close()
	return longLink
}
