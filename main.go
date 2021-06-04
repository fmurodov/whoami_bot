package main

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	var botToken string = os.Getenv("BOT_TOKEN")
	var text string
	if len(botToken) == 0 {
		log.Fatal("Port is not set")
	}

	log.Println("Starting bot...")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", strconv.Itoa(update.Message.From.ID), update.Message.Text)

		if update.Message.Text == "/me" || update.Message.Text == "/whoami" {
			text = ("chat id: " + strconv.Itoa(update.Message.From.ID) +
				"\nusername: " + update.Message.From.UserName +
				"\nname: " + update.Message.From.FirstName + " " + update.Message.From.LastName)

		} else if update.Message.Text == "/start" {
			text = "Hi! I know all about you!\nSend me /whoami command"

		} else {
			text = "send /me or /whoami for get info"

		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
