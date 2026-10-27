package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Every request takes one second, like a slow API.
func main() {
	http.HandleFunc("/page/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second)
		fmt.Fprintf(w, "%s ok", r.URL.Path)
	})
	log.Fatal(http.ListenAndServe("localhost:9090", nil))
}
