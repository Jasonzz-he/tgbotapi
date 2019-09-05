package tgbotapi

import "github.com/Jasonzz-he/tgbotapi/botproto"

func (b *Bot) SendGame(in *botproto.SendGame) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("sendGame", in, rst)
	return rst, err
}

func SendGame(in *botproto.SendGame) (*botproto.Message, error) {
	return globalBot.SendGame(in)
}

func (b *Bot) SetGameScore(in *botproto.SetGameScore) error {
	err := b.Post("setGameScore", in, nil)
	return err
}

func SetGameScore(in *botproto.SetGameScore) error {
	return globalBot.SetGameScore(in)
}

func (b *Bot) GetGameHighScores(in *botproto.GetGameHighScores) ([]*botproto.GameHighScore, error) {
	var rst = make([]*botproto.GameHighScore, 0)
	err := b.Post("getGameHighScores", in, rst)
	return rst, err
}

func GetGameHighScores(in *botproto.GetGameHighScores) ([]*botproto.GameHighScore, error) {
	return globalBot.GetGameHighScores(in)
}
