package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/tucnak/telebot.v2"
)

var tgBot *telebot.Bot = nil
var tgUserId int

func initTelegram(twitterApi *twitter.Client) {
	log.Println("Get config")
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	channelId, _ := strconv.ParseInt(os.Getenv("TEST_CHANNEL_ID"), 10, 64)
	tgUserId, _ = strconv.Atoi(os.Getenv("TELEGRAM_USER_ID"))

	log.Println("Start bot")
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		panic(err)
	}

	tgBot = bot

	bot.Handle(telebot.OnChannelPost, func(inputMessage *telebot.Message) {

		if inputMessage.Chat.ID != channelId {
			log.Println(inputMessage.Chat.ID, " not accepted")
			return
		}

		if !strings.Contains(inputMessage.Text, "#twitter") {
			log.Println("No #twitter tag!")
			return
		}

		tweetText := strings.ReplaceAll(inputMessage.Text, "#twitter", "")
		log.Println(tweetText)
		// Send a Tweet
		_, _, err := twitterApi.Statuses.Update(tweetText, nil)

		if err != nil {
			tgLog(err.Error(), tgUserId)
		} else {
			log.Println("Tweet has been created!")
		}
	})

	fmt.Println(getChannelInfo(channelId, bot))
	log.Println("Waiting any channel post")

	bot.Start()
}

func getChannelInfo(channelID int64, bot *telebot.Bot) string {
	chat, err := bot.ChatByID(fmt.Sprint(channelID))

	if err != nil {
		tgLog(err.Error(), tgUserId)
		return ""
	}

	return fmt.Sprintf(
		"Channel: %s\nUserName: %s\nDescription: %s",
		chat.Title, chat.Username, chat.Description,
	)
}
