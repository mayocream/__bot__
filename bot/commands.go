package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Commands ...
func (b *Bot) Commands(api *tgbotapi.BotAPI, update *tgbotapi.Update) error {
	switch update.Message.Command() {
	case "ecs_instances":
		resp, err := b.AliCloudECS.GetInstances()
		if err != nil {
			return err
		}
		for _, ins := range resp.Instances.Instance {
			
		}
	default:
	}
	return nil
}
