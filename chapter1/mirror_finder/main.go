package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rprajapati0067/hands-on-restful-api-golang/chapter1/mirrors"
)

type response struct {
	FastestURL string        `json:"fastest_url"`
	Latency    time.Duration `json:"latency"`
}

func findFastest(urls []string) response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)

	for _, url := range urls {
		mirrorURL := url
		go func() {
			start := time.Now()
			_, err := http.Get(mirrorURL + "/README")
			latency := time.Now().Sub(start) / time.Millisecond
			if err == nil {
				urlChan <- mirrorURL
				latencyChan <- latency
			}
		}()

	}
	return response{<-urlChan, <-latencyChan}
}

func main() {
	http.HandleFunc("/fastest_url", func(w http.ResponseWriter, r *http.Request) {
		response := findFastest(mirrors.MirrorList)
		respJSON, _ := json.Marshal(response)
		fmt.Println(string(respJSON))
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJSON)
	})
	port := ":1323"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Starting server on port %s\n ", port)
	log.Fatal(server.ListenAndServe())

}
