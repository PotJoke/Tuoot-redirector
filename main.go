package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	check()
	fmt.Print("Check completed\n")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ok")
	})
	http.HandleFunc("/test", PostGetTest)

	log.Fatal(http.ListenAndServe(port, nil))
}

func PostGetTest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		io.WriteString(w, "POST METHOD")
	case "GET":
		io.WriteString(w, "GET METHOD")
	default:
		io.WriteString(w, "METHOD NOT ALLOWED")
	}
}
