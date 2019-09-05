package tgbotapi

import "github.com/Jasonzz-he/tgbotapi/botproto"

func (b *Bot) SendSticker(in *botproto.SendSticker) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("sendSticker", in, rst)
	return rst, err
}

func SendSticker(in *botproto.SendSticker) (*botproto.Message, error) {
	return globalBot.SendSticker(in)
}

func (b *Bot) GetStickerSet(in *botproto.GetStickerSet) error {
	err := b.Post("getStickerSet", in, nil)
	return err
}

func GetStickerSet(in *botproto.GetStickerSet) error {
	return globalBot.GetStickerSet(in)
}

func (b *Bot) UploadStickerFile(in *botproto.UploadStickerFile) (*botproto.File, error) {
	var rst = new(botproto.File)
	err := b.Post("uploadStickerFile", in, rst)
	return rst, err
}

func UploadStickerFile(in *botproto.UploadStickerFile) (*botproto.File, error) {
	return globalBot.UploadStickerFile(in)
}

func (b *Bot) CreateNewStickerSet(in *botproto.CreateNewStickerSet) (*bool, error) {
	var rst = new(bool)
	err := b.Post("createNewStickerSet", in, rst)
	return rst, err
}

func CreateNewStickerSet(in *botproto.CreateNewStickerSet) (*bool, error) {
	return globalBot.CreateNewStickerSet(in)
}

func (b *Bot) SetStickerPositionInSet(in *botproto.SetStickerPositionInSet) (*bool, error) {
	var rst = new(bool)
	err := b.Post("setStickerPositionInSet", in, rst)
	return rst, err
}

func SetStickerPositionInSet(in *botproto.SetStickerPositionInSet) (*bool, error) {
	return globalBot.SetStickerPositionInSet(in)
}

func (b *Bot) DeleteStickerFromSet(in *botproto.DeleteStickerFromSet) (*bool, error) {
	var rst = new(bool)
	err := b.Post("deleteStickerFromSet", in, rst)
	return rst, err
}

func DeleteStickerFromSet(in *botproto.DeleteStickerFromSet) (*bool, error) {
	return globalBot.DeleteStickerFromSet(in)
}
