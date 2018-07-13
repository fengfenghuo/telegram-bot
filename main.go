package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	botApi := "581459553:AAGziS8gfuhXTOp6cJz8U12rSRgJiuGUz8Q"
	// test 524361350:AAGneffMyxXPDNOZB5hOGvCP198rMv5PuG0
	bot, err := tgbotapi.NewBotAPI(botApi)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if update.Message.NewChatMembers != nil {
			str := "Welcome "
			for _, temp := range *update.Message.NewChatMembers {
				str += temp.FirstName + " " + temp.LastName + " "
			}
			str += "join group!!!"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, str)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)

			deletemsg := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
			bot.Send(deletemsg)
		} else if update.Message.LeftChatMember != nil {
			deletemsg := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
			bot.Send(deletemsg)
		}
	}
}
