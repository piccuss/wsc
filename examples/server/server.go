package main

import (
	"log"
	"net/http"

	core "github.com/piccuss/wsc"
)

func main() {
	http.HandleFunc("/hello", core.HandlerEcho)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
