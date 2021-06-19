package main

import (
	"net/http"
	"os"
)

func PingHandler(resp http.ResponseWriter, _ *http.Request) {
	resp.Write([]byte("Hi there! I'm BaLiK_ShitPoster bot!"))
}

func main() {
	twitterApi := initTwitter()
	initTelegram(twitterApi)

	http.HandleFunc("/", PingHandler)
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
