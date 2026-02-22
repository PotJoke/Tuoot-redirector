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

	log.Fatal(http.ListenAndServe(port, nil))
}
