package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "处理器: %s\n", "indexHandler")
	if err != nil {
		return
	}

}
func hiHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "处理器: %s\n", "hiHandler")
	if err != nil {
		return
	}

}
func webHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "处理器: %s\n", "webHandler")
	if err != nil {
		return
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/hi", hiHandler)
	mux.HandleFunc("/hi/web", webHandler)

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
