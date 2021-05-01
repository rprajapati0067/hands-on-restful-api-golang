package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got Parameter catagery: %s!\n", queryParams["category"][0])
	fmt.Fprintf(w, "Got Parameter id: %v\n", queryParams["id"][0])

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles", QueryHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
