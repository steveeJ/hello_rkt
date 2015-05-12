package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Not enough arguments.")
		os.Exit(1)
	}
	root_dir := os.Args[1]
	port := ":" + os.Args[2]
	fmt.Printf("Serving files from %v on port %v", root_dir, port)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(root_dir))))
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
