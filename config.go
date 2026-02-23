package main

import (
	"fmt"
	"math/rand"
	"os"
)

var port = os.Getenv("PORT")
var certfile_path = os.Getenv("CERTFILE_PATH")
var keyfile_path = os.Getenv("KEYFILE_PATH")

func envcheck() {
	if port == "" {
		port = ":8080"
		fmt.Printf("PORT not set, defaulting to %s\n", port)
	}
	if certfile_path == "" && keyfile_path == "" {
		fmt.Println("CERTFILE_PATH and KEYFILE_PATH not set")
		return
	}
}

func generateUserAgent() string {
	userAgents := []string{
		"Dalvik/2.1.0 (Linux; U; Android 8.0.0; SM-J337A Build/R16NW)",
		"Dalvik/2.1.0 (Linux; U; Android 9; Mediatek MT8173 Chromebook Build/R93-14092.57.0)",
		"Dalvik/2.1.0 (Linux; U; Android 11; PDEM30 Build/RKQ1.200710.002)",
		"Mozilla/5.0 (Linux; Android 8.1.0; Armor_3 Build/O11019; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/69.0.3497.100 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 10.0; P40 Pro Build/LMY47I; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/50.0.2661.86 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 6.0; U FEEL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Mobile Safari/537.36 OPR/58.2.2878.53403",
		"Mozilla/5.0 (Linux; Android 9; LM-G710VM) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.101 Mobile Safari/537.36",
		"Dalvik/2.1.0 (Linux; U; Android 11; SM-F711U Build/RP1A.200720.012)",
		"Dalvik/1.6.0 (Linux; U; Android 4.2.1; PTT826D Build/JB)",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 YaBrowser/21.8.1.476 Yowser/2.5 Safari/537.36",
	}
	return userAgents[rand.Intn(len(userAgents))]
}
