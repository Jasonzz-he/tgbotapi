package tgbotapi

import (
	"github.com/Jasonzz-he/tgbotapi/botproto"
	"github.com/smartystreets/assertions"
	"log"
	"testing"
)

func TestBot_EditMessageText(t *testing.T) {
	err := EditMessageText(&botproto.EditMessageText{
		ChatId:                "",    // string
		MessageId:             0,     // int32
		InlineMessageId:       "",    // string
		Text:                  "",    // string
		ParseMode:             "",    // string
		DisableWebPagePreview: false, // bool
		ReplyMarkup:           nil,   // InlineKeyboardMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
}

func TestBot_EditMessageCaption(t *testing.T) {
	err := EditMessageCaption(&botproto.EditMessageCaption{
		ChatId:          "",  // string
		MessageId:       0,   // int32
		InlineMessageId: "",  // string
		Caption:         "",  // string
		ParseMode:       "",  // string
		ReplyMarkup:     nil, // InlineKeyboardMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
}

func TestBot_EditMessageMedia(t *testing.T) {
	err := EditMessageMedia(&botproto.EditMessageMedia{
		ChatId:          "",  // string
		MessageId:       0,   // int32
		InlineMessageId: "",  // string
		Media:           nil, // InputMedia
		ReplyMarkup:     nil, // InlineKeyboardMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
}

func TestBot_EditMessageReplyMarkup(t *testing.T) {
	err := EditMessageReplyMarkup(&botproto.EditMessageReplyMarkup{
		ChatId:          "",  // string
		MessageId:       0,   // int32
		InlineMessageId: "",  // string
		ReplyMarkup:     nil, // InlineKeyboardMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
}

func TestBot_StopPoll(t *testing.T) {
	err := StopPoll(&botproto.StopPoll{
		ChatId:      "",  // string
		MessageId:   0,   // int32
		ReplyMarkup: nil, // InlineKeyboardMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
}

func TestBot_DeleteMessage(t *testing.T) {
	rst, err := DeleteMessage(&botproto.DeleteMessage{
		ChatId:    "", // string
		MessageId: 0,  // int32
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}
