package tgbotapi

import "github.com/Jasonzz-he/tgbotapi/botproto"

func (b *Bot) EditMessageText(in *botproto.EditMessageText) error {
	err := b.Post("editMessageText", in, nil)
	return err
}

func EditMessageText(in *botproto.EditMessageText) error {
	return globalBot.EditMessageText(in)
}

func (b *Bot) EditMessageCaption(in *botproto.EditMessageCaption) error {
	err := b.Post("editMessageCaption", in, nil)
	return err
}

func EditMessageCaption(in *botproto.EditMessageCaption) error {
	return globalBot.EditMessageCaption(in)
}

func (b *Bot) EditMessageMedia(in *botproto.EditMessageMedia) error {
	err := b.Post("editMessageMedia", in, nil)
	return err
}

func EditMessageMedia(in *botproto.EditMessageMedia) error {
	return globalBot.EditMessageMedia(in)
}

func (b *Bot) EditMessageReplyMarkup(in *botproto.EditMessageReplyMarkup) error {
	err := b.Post("editMessageReplyMarkup", in, nil)
	return err
}

func EditMessageReplyMarkup(in *botproto.EditMessageReplyMarkup) error {
	return globalBot.EditMessageReplyMarkup(in)
}

func (b *Bot) StopPoll(in *botproto.StopPoll) error {
	err := b.Post("stopPoll", in, nil)
	return err
}

func StopPoll(in *botproto.StopPoll) error {
	return globalBot.StopPoll(in)
}

func (b *Bot) DeleteMessage(in *botproto.DeleteMessage) (*bool, error) {
	var rst = new(bool)
	err := b.Post("deleteMessage", in, rst)
	return rst, err
}

func DeleteMessage(in *botproto.DeleteMessage) (*bool, error) {
	return globalBot.DeleteMessage(in)
}
