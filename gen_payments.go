package tgbotapi

import "github.com/Jasonzz-he/tgbotapi/botproto"

func (b *Bot) SendInvoice(in *botproto.SendInvoice) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("sendInvoice", in, rst)
	return rst, err
}

func SendInvoice(in *botproto.SendInvoice) (*botproto.Message, error) {
	return globalBot.SendInvoice(in)
}
