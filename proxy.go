package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", handleReq)
	log.Println("CORS Proxy running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	// Disallow all methods except GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse URL from query string
	target := r.URL.Query().Get("url")
	if target == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Verify URL format
	verified, err := url.ParseRequestURI(target)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Perform GET request
	resp, err := http.Get(verified.String())
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Permissive CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Method", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(resp.StatusCode)

	// Forward response
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
