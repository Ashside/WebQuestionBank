package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there!")

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hi)

	myServer := &http.Server{
		Addr:         ":8081",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := myServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
