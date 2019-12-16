package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/short", handleShort)
	http.HandleFunc("/long", handleLong)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
