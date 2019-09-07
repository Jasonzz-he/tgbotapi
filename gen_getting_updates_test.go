package tgbotapi

import (
	"github.com/Jasonzz-he/tgbotapi/botproto"
	"github.com/smartystreets/assertions"
	"log"
	"testing"
)

func TestBot_GetUpdates(t *testing.T) {
	err := GetUpdates(&botproto.GetUpdates{
		Offset:         0,   // int32
		Limit:          0,   // int32
		Timeout:        0,   // int32
		AllowedUpdates: nil, // repeated string
	})
	log.Println(assertions.ShouldBeNil(err))
}

func TestBot_SetWebhook(t *testing.T) {
	rst, err := SetWebhook(&botproto.SetWebhook{
		Url:            "",  // string
		Certificate:    nil, // InputFile
		MaxConnections: 0,   // int32
		AllowedUpdates: nil, // repeated string
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_DeleteWebhook(t *testing.T) {
	rst, err := DeleteWebhook()
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_GetWebhookInfo(t *testing.T) {
	err := GetWebhookInfo()
	log.Println(assertions.ShouldBeNil(err))
}
