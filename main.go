package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	check()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ok")
	})
	http.HandleFunc("/redirect", redirector)

	log.Fatal(http.ListenAndServe(port, nil))
}

func redirector(w http.ResponseWriter, r *http.Request) {

	if r.URL.RawQuery == "" {
		io.WriteString(w, "REDIRECTION ERROR: URL EMPTY")
		return
	}

	client := &http.Client{}
	request := &http.Request{}

	body_resp := []byte{}

	switch r.Method {
	case "GET":
		request, _ = http.NewRequest("GET", r.URL.RawQuery, r.Body)
	case "POST":
		request, _ = http.NewRequest("POST", r.URL.RawQuery, r.Body)

	default:
		io.WriteString(w, "METHOD NOT ALLOWED")
	}

	request.Header.Set("Cookie", r.Header.Get("Cookie"))

	resp, _ := client.Do(request)
	body_resp, _ = io.ReadAll(resp.Body)

	io.Writer.Write(w, body_resp)
}
