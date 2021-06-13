package main

import (
	"log"
	"time"

	"gopkg.in/tucnak/telebot.v2"
)

func tgLog(text string, id int) {
	log.Println(text)
	if tgBot != nil {
		tgBot.Send(&telebot.User{ID: id}, time.Now().String()+" : "+text)
	}
}
