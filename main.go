package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	envcheck()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ok")
	})
	http.HandleFunc("/redirect", redirector)

	log.Fatal(http.ListenAndServeTLS(port, certfile_path, keyfile_path, nil))
}

func redirector(w http.ResponseWriter, r *http.Request) {

	if r.URL.RawQuery == "" {
		io.WriteString(w, "REDIRECTION ERROR: URL EMPTY")
		return
	}

	client := &http.Client{}
	body_resp := []byte{}

	request, _ := http.NewRequest(r.Method, r.URL.RawQuery, r.Body)
	request.Header.Set("User-Agent", generateUserAgent())
	request.Header.Set("Cookie", r.Header.Get("Cookie"))

	resp, _ := client.Do(request)
	body_resp, _ = io.ReadAll(resp.Body)

	io.Writer.Write(w, body_resp)
}
