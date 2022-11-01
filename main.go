package main

import (
	"freeEda/burger"
	"freeEda/config"
	"freeEda/wkusno"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config := config.Load()
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		switch update.Message.Text {
		case "/start":
			message := config.HelloText
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			bot.Send(msg)
		case config.BurgerCommand:
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			cupoons, err := burger.GetBurgerCupoons(config.ApiBurger)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "ошибка")
				bot.Send(msg)
			}
			for i := 0; i < len(cupoons.Response.Dishes); i++ {
				message := "Акция: " + cupoons.Response.Dishes[i].Name +
					"\nЦена: " + strconv.Itoa(cupoons.Response.Dishes[i].Price/100) +
					"\nОписание: " + cupoons.Response.Dishes[i].Description +
					"\nКод: " + cupoons.Response.Dishes[i].Code
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				url := tgbotapi.NewInputMediaPhoto(tgbotapi.FileURL(cupoons.Response.Dishes[i].Image))
				mediaGroup := tgbotapi.NewMediaGroup(update.Message.Chat.ID, []interface{}{
					url,
				})
				bot.SendMediaGroup(mediaGroup)
				bot.Send(msg)
			}
		case config.WkusnoCommand:
			cupoons, err := wkusno.GetWkusnoCupoons(config.ApiWkusno)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "ошибка")
				bot.Send(msg)
			}
			for i := 0; i < len(cupoons.Items); i++ {
				message := "Акция: " + cupoons.Items[i].Name
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				url := tgbotapi.NewInputMediaPhoto(tgbotapi.FileURL("https://vkusnoitochka.ru/" + cupoons.Items[i].Pic))
				mediaGroup := tgbotapi.NewMediaGroup(update.Message.Chat.ID, []interface{}{
					url,
				})
				bot.SendMediaGroup(mediaGroup)
				bot.Send(msg)
			}
		}
	}
}
