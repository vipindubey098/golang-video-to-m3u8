package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("../output"))

	http.Handle("/", fs)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
