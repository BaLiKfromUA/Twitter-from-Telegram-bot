package main

import (
	"net/http"
	"os"
)

func main() {
	twitterApi := initTwitter()
	initTelegram(twitterApi)
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
