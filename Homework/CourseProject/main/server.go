package main

import (
	"fmt"
	"net/http"
)

// handle short link request
// get request
func handleShort(response http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	shortLink := request.Form["short_link"]
	url := shortLinkToLongLink(shortLink[0])
	http.Redirect(response, request, url, http.StatusFound)
}

// handle long link request
// post request
func handleLong(response http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	longLink := request.Form["long_link"]
	fmt.Println("longLink: ", longLink)
	_, _ = fmt.Fprintf(response, longLinkToShortLink(longLink[0]))
}
