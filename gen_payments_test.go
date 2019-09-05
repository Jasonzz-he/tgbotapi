package tgbotapi

import (
	"github.com/Jasonzz-he/tgbotapi/botproto"
	"github.com/smartystreets/assertions"
	"log"
	"testing"
)

func TestBot_SendInvoice(t *testing.T) {
	rst, err := SendInvoice(&botproto.SendInvoice{
		ChatId:                    0,     // int32
		Title:                     "",    // string
		Description:               "",    // string
		Payload:                   "",    // string
		ProviderToken:             "",    // string
		StartParameter:            "",    // string
		Currency:                  "",    // string
		Prices:                    nil,   // repeated LabeledPrice
		ProviderData:              "",    // string
		PhotoUrl:                  "",    // string
		PhotoSize:                 0,     // int32
		PhotoWidth:                0,     // int32
		PhotoHeight:               0,     // int32
		NeedName:                  false, // bool
		NeedPhoneNumber:           false, // bool
		NeedEmail:                 false, // bool
		NeedShippingAddress:       false, // bool
		SendPhoneNumberToProvider: false, // bool
		SendEmailToProvider:       false, // bool
		IsFlexible:                false, // bool
		DisableNotification:       false, // bool
		ReplyToMessageId:          0,     // int32
		ReplyMarkup:               nil,   // InlineKeyboardMarkup
	})
	assertions.ShouldBeNil(err)
	log.Printf("%#v", rst)
}
