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
	http.HandleFunc("/redirect", redirector)

	log.Fatal(http.ListenAndServe(port, nil))
}

func redirector(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.RawQuery == "" {
			io.WriteString(w, "GET METHOD")
			return
		}

		client := &http.Client{}

		request, _ := http.NewRequest("POST", r.URL.RawQuery, nil)
		request.Header.Set("Cookie", r.Header.Get("Cookie"))
		resp, _ := client.Do(request)

		body, _ := io.ReadAll(resp.Body)
		io.Writer.Write(w, body)

	default:
		io.WriteString(w, "METHOD NOT ALLOWED")
	}
}
