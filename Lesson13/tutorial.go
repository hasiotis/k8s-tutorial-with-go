package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const DumpFile = "/data/dump.txt"

var special, remote, secret string

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, golang (special: %s, secret: %s)!", special, secret)
	log.Printf("Root request from %s", r.RemoteAddr)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	value, err := ioutil.ReadFile(DumpFile)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprintf(w, "Value is [%s]", value)
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	value := strings.Replace(r.URL.Path, "/set/", "", 1)
	err := ioutil.WriteFile(DumpFile, []byte(value), 0644)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprintf(w, "Value [%s] set", value)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	err := os.Truncate(DumpFile, 0)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprintf(w, "Truncated dump file")
}

func ready(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
	log.Printf("Ready")
}

func main() {
	special = os.Getenv("TUTORIAL_SPECIAL")
	remote = os.Getenv("TUTORIAL_REMOTE")
	secret = os.Getenv("TUTORIAL_SECRET")

	os.OpenFile(DumpFile, os.O_RDONLY|os.O_CREATE, 0666)

	http.HandleFunc("/", handler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/set/", setHandler)
	http.HandleFunc("/reset", resetHandler)
	http.HandleFunc("/ready", ready)

	http.ListenAndServe(":8000", nil)
}
