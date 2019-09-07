package tgbotapi

import "github.com/Jasonzz-he/tgbotapi/botproto"

func (b *Bot) GetUpdates(in *botproto.GetUpdates) error {
	err := b.Post("getUpdates", in, nil)
	return err
}

func GetUpdates(in *botproto.GetUpdates) error {
	return globalBot.GetUpdates(in)
}

func (b *Bot) SetWebhook(in *botproto.SetWebhook) (*bool, error) {
	var rst = new(bool)
	err := b.Post("setWebhook", in, rst)
	return rst, err
}

func SetWebhook(in *botproto.SetWebhook) (*bool, error) {
	return globalBot.SetWebhook(in)
}

func (b *Bot) DeleteWebhook() (*bool, error) {
	var rst = new(bool)
	err := b.Post("deleteWebhook", nil, rst)
	return rst, err
}

func DeleteWebhook() (*bool, error) {
	return globalBot.DeleteWebhook()
}

func (b *Bot) GetWebhookInfo() error {
	err := b.Post("getWebhookInfo", nil, nil)
	return err
}

func GetWebhookInfo() error {
	return globalBot.GetWebhookInfo()
}
