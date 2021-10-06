package telegram

import (
	"fmt"
	"github.com/Psh777/visa-backend/modules/config"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
	"time"
)

var bot *tgbotapi.BotAPI

func Init(myConfig config.Env) {

	for {

		var err error
		bot, err = tgbotapi.NewBotAPI(myConfig.TelegramBot)

		if err != nil {
			fmt.Println("Telegram connection refused", err)
			time.Sleep(time.Second * 10)
		} else {
			break
		}
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		fmt.Println("Telegram error:", err)
		return
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//if update.Message.ReplyToMessage.Text

		switch strings.ToLower(update.Message.Text) {

		case "get id":
			{
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint(update.Message.Chat.ID))
				_, _ = bot.Send(msg)
			}

		case "hi":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello")
			_, _ = bot.Send(msg)

		case "help":
			var mess string
			mess = mess + "hi - check connect bot\n"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, mess)
			_, _ = bot.Send(msg)
		}

	}

}

func SendMsgBot(text string) {
	var list [3]int64
	list[0] = 7314332
	list[1] = 3221264
	list[2] = 269490674
	//list[1] = 2
	for i := 0; i < len(list); i++ {
		msg := tgbotapi.NewMessage(list[i], text)
		_, err := bot.Send(msg)
		if err != nil {
			fmt.Println(err)
		}

	}
}
