package main

import (
	"log"
	"net/http"
)

func main() {
    port := ":5000"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request from %v\n", r.RemoteAddr)
		w.Write([]byte("Hello from App Container"))
	})

    log.Printf("About to listen on port %v\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
