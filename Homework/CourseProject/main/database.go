package main

func longLinkToShortLink(longLink string) string {
	var shortLink string
	shortLink = "I am a short link"
	//db, _ := sql.Open("mysql", "root:root@tcp(47.102.205.224:3306)/test?charset=utf8")
	//shortLink = longLink
	//_, _ = db.Query("insert into link (short_link, long_link) values (?, ?);", shortLink, longLink)
	//_ = db.Close()
	return shortLink
}

func shortLinkToLongLink(shortLink string) string {
	var longLink string
	longLink = "I am a long link"
	//linkIndex := shortLinkToLongInt(shortLink)
	//db, _ := sql.Open("mysql", "root:root@tcp(47.102.205.224:3306)/test?charset=utf8")
	//row:= db.QueryRow("select long_link from link where link_index = ?;", linkIndex)
	//_ = row.Scan(&longLink)
	//_ = db.Close()
	return longLink
}