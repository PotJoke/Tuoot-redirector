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

	fmt.Print(http.Get("http://localhost" + port + "/test"))
}

func PostGetTest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		io.WriteString(w, "POST METHOD")
	case "GET":
		if r.URL.RawQuery == "" {
			io.WriteString(w, "GET METHOD")
			return
		}
		res, _ := http.Get(r.URL.RawQuery)
		body, _ := io.ReadAll(res.Body)
		io.WriteString(w, string(body))

	default:
		io.WriteString(w, "METHOD NOT ALLOWED")
	}
}
