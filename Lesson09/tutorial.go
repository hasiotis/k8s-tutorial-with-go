package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var special, remote, secret string

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, golang (special: %s, secret: %s)!", special, secret)
	log.Printf("Root request from %s", r.RemoteAddr)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(remote)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	w.Write(body)
	log.Printf("Get request from %s", r.RemoteAddr)
}

func ready(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
	log.Printf("Ready")
}

func main() {
	special = os.Getenv("TUTORIAL_SPECIAL")
	remote = os.Getenv("TUTORIAL_REMOTE")
	secret = os.Getenv("TUTORIAL_SECRET")

	http.HandleFunc("/", handler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/ready", ready)

	time.Sleep(10 * time.Second)
	http.ListenAndServe(":8000", nil)
}
