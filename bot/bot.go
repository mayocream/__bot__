package bot

import (
	"bot/pkg/alicloud"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Bot ...
type Bot struct {
	API         *tgbotapi.BotAPI
	AliCloudECS *alicloud.ECSClient
	Config      *Config
}

// NewBot ...
func NewBot(config *Config) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(config.APIToken)
	if err != nil {
		return nil, err
	}

	ecs, err := alicloud.NewECSClient(&config.AliCloud)
	if err != nil {
		return nil, err
	}

	b := &Bot{
		API:         api,
		AliCloudECS: ecs,
	}
	return b, nil
}

// Run ...
func (b *Bot) Run() error {
	log.Printf("Authorized on account %s", b.API.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updateChan := b.API.GetUpdatesChan(u)

	for update := range updateChan {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message.From.UserName != b.Config.MasterUsername {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		b.API.Send(msg)
	}

	return nil
}
