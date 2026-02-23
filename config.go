package main

import (
	"fmt"
	"os"
)

var port = os.Getenv("PORT")
var certfile_path = os.Getenv("CERTFILE_PATH")
var keyfile_path = os.Getenv("KEYFILE_PATH")

func check() {
	if port == "" {
		port = ":8080"
		fmt.Printf("PORT not set, defaulting to %s\n", port)
	}
	if certfile_path == "" && keyfile_path == "" {
		fmt.Println("CERTFILE_PATH and KEYFILE_PATH not set")
		return
	}
}
