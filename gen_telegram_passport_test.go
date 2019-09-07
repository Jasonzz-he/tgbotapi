package tgbotapi

import (
	"github.com/Jasonzz-he/tgbotapi/botproto"
	"github.com/smartystreets/assertions"
	"log"
	"testing"
)

func TestBot_SetPassportDataErrors(t *testing.T) {
	rst, err := SetPassportDataErrors(&botproto.SetPassportDataErrors{
		UserId: 0,   // int32
		Errors: nil, // repeated PassportElementError
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}
