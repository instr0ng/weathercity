package main

import (
	"log"
	"math"
	"strconv"

	owm "github.com/briandowns/openweathermap"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Currentpogoda(city string) string {
	w, err := owm.NewCurrent("C", "ru", "eb02daa1e670095dd12a65ebd52a2fd1")
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName(city)
	if w.Name != "" {
		return "Сейчас " + strconv.FormatFloat(math.Round(w.Main.Temp), 'f', 0, 64) + "°C " + w.Weather[0].Description
	}
	return "Некорректный город"
}

func telegrambot() {
	bot, err := tgbotapi.NewBotAPI("5504867100:AAG4hdvRYmGGCXjHl-J1CZmuxI23FfCR0Ws")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, я погодный бот. Введи город и узнаешь какая в нём погода!")
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, Currentpogoda(update.Message.Text))
				bot.Send(msg)
			}
		}
	}
}

func main() {
	telegrambot()
}
