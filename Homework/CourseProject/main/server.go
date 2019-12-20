package main

import (
	"fmt"
	"net/http"
)

// handle short link request
func handleShort(response http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	shortLink := request.Form["short_link"]
	fmt.Println("shortLink: ", shortLink)
	_, _ = fmt.Fprintf(response, shortLinkToLongLink(shortLink[0]))
}

// handle long link request
func handleLong(response http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	longLink := request.Form["long_link"]
	fmt.Println("longLink: ", longLink)
	_, _ = fmt.Fprintf(response, longLinkToShortLink(longLink[0]))
}
