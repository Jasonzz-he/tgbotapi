package tgbotapi

import "github.com/Jasonzz-he/tgbotapi/botproto"

func (b *Bot) SetPassportDataErrors(in *botproto.SetPassportDataErrors) (*bool, error) {
	var rst = new(bool)
	err := b.Post("setPassportDataErrors", in, rst)
	return rst, err
}

func SetPassportDataErrors(in *botproto.SetPassportDataErrors) (*bool, error) {
	return globalBot.SetPassportDataErrors(in)
}
