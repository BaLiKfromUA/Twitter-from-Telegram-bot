package main

func main() {
	twitterApi := initTwitter()
	initTelegram(twitterApi)
}
