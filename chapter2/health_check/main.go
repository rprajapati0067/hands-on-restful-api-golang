package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func HeathCheck(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	io.WriteString(w, currentTime.String())

}

func main() {
	http.HandleFunc("/health", HeathCheck)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
