package tgbotapi

import (
	"github.com/Jasonzz-he/tgbotapi/botproto"
	"github.com/smartystreets/assertions"
	"log"
	"testing"
)

func TestBot_SendSticker(t *testing.T) {
	rst, err := SendSticker(&botproto.SendSticker{
		ChatId:              0,     // int32
		Sticker:             "",    // string
		DisableNotification: false, // bool
		ReplyToMessageId:    0,     // int32
		ReplyMarkup:         nil,   // ReplyMarkup
	})
	assertions.ShouldBeNil(err)
	log.Printf("%#v", rst)
}

func TestBot_GetStickerSet(t *testing.T) {
	err := GetStickerSet(&botproto.GetStickerSet{
		Name: "", // string
	})
	assertions.ShouldBeNil(err)
}

func TestBot_UploadStickerFile(t *testing.T) {
	rst, err := UploadStickerFile(&botproto.UploadStickerFile{
		UserId:     0,   // int32
		PngSticker: nil, // InputFile
	})
	assertions.ShouldBeNil(err)
	log.Printf("%#v", rst)
}

func TestBot_CreateNewStickerSet(t *testing.T) {
	rst, err := CreateNewStickerSet(&botproto.CreateNewStickerSet{
		UserId:        0,     // int32
		Name:          "",    // string
		Title:         "",    // string
		PngSticker:    "",    // string
		Emojis:        "",    // string
		ContainsMasks: false, // bool
		MaskPosition:  nil,   // MaskPosition
	})
	assertions.ShouldBeNil(err)
	log.Printf("%#v", rst)
}

func TestBot_SetStickerPositionInSet(t *testing.T) {
	rst, err := SetStickerPositionInSet(&botproto.SetStickerPositionInSet{
		Sticker:  "", // string
		Position: 0,  // int32
	})
	assertions.ShouldBeNil(err)
	log.Printf("%#v", rst)
}

func TestBot_DeleteStickerFromSet(t *testing.T) {
	rst, err := DeleteStickerFromSet(&botproto.DeleteStickerFromSet{
		Sticker: "", // string
	})
	assertions.ShouldBeNil(err)
	log.Printf("%#v", rst)
}
