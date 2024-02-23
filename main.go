package main

import (
	"fmt"
	"net/http"
	"time"
)

func measureLoadTime(url string) (time.Duration, error) {
	start := time.Now()
	_, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	elapsed := time.Since(start)
	return elapsed, nil
}

func loadTimeHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}
	elapsedTime, err := measureLoadTime(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error measuring load time: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "The URL %s loaded in %s", url, elapsedTime)
}

func main() {
	http.HandleFunc("/loadtime", loadTimeHandler)
	http.ListenAndServe(":8080", nil)
}
