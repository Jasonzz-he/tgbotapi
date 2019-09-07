package tgbotapi

import (
	"github.com/Jasonzz-he/tgbotapi/botproto"
	"github.com/smartystreets/assertions"
	"log"
	"testing"
)

func TestBot_SendGame(t *testing.T) {
	rst, err := SendGame(&botproto.SendGame{
		ChatId:              0,     // int32
		GameShortName:       "",    // string
		DisableNotification: false, // bool
		ReplyToMessageId:    0,     // int32
		ReplyMarkup:         nil,   // InlineKeyboardMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SetGameScore(t *testing.T) {
	err := SetGameScore(&botproto.SetGameScore{
		UserId:             0,     // int32
		Score:              0,     // int32
		Force:              false, // bool
		DisableEditMessage: false, // bool
		ChatId:             0,     // int32
		MessageId:          0,     // int32
		InlineMessageId:    "",    // string
	})
	log.Println(assertions.ShouldBeNil(err))
}

func TestBot_GetGameHighScores(t *testing.T) {
	rst, err := GetGameHighScores(&botproto.GetGameHighScores{
		UserId:          0,  // int32
		ChatId:          0,  // int32
		MessageId:       0,  // int32
		InlineMessageId: "", // string
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}
