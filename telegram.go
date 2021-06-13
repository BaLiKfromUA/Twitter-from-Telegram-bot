package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/tucnak/telebot.v2"
)

var tgBot *telebot.Bot = nil
var tgUserId int

func initTelegram() {
	log.Println("Get config")
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	channelId, _ := strconv.ParseInt(os.Getenv("TELEGRAM_CHANNEL_ID"), 10, 64)
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
		log.Println(inputMessage.Chat.ID)

		if inputMessage.Chat.ID != channelId {
			log.Println(inputMessage.Chat.ID, " not accepted")
			return
		}

		fmt.Println(inputMessage)
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
