package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var special string

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, golang (special: %s)!", special)
	log.Printf("Request from %s", r.RemoteAddr)
}

func main() {
	special = os.Getenv("TUTORIAL_SPECIAL")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
