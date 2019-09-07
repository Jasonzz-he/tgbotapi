package tgbotapi

import (
	"github.com/Jasonzz-he/tgbotapi/botproto"
	"github.com/smartystreets/assertions"
	"log"
	"testing"
)

func TestBot_GetMe(t *testing.T) {
	rst, err := GetMe()
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SendMessage(t *testing.T) {
	rst, err := SendMessage(&botproto.SendMessage{
		ChatId:                "",    // string
		Text:                  "",    // string
		ParseMode:             "",    // string
		DisableWebPagePreview: false, // bool
		DisableNotification:   false, // bool
		ReplyToMessageId:      0,     // int32
		ReplyMarkup:           nil,   // ReplyMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_ForwardMessage(t *testing.T) {
	rst, err := ForwardMessage(&botproto.ForwardMessage{
		ChatId:              "",    // string
		FromChatId:          "",    // string
		DisableNotification: false, // bool
		MessageId:           0,     // int32
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SendPhoto(t *testing.T) {
	rst, err := SendPhoto(&botproto.SendPhoto{
		ChatId:              "",    // string
		Photo:               "",    // string
		Caption:             "",    // string
		ParseMode:           "",    // string
		DisableNotification: false, // bool
		ReplyToMessageId:    0,     // int32
		ReplyMarkup:         nil,   // ReplyMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SendAudio(t *testing.T) {
	rst, err := SendAudio(&botproto.SendAudio{
		ChatId:              "",    // string
		Audio:               "",    // string
		Caption:             "",    // string
		ParseMode:           "",    // string
		Duration:            0,     // int32
		Performer:           "",    // string
		Title:               "",    // string
		Thumb:               "",    // string
		DisableNotification: false, // bool
		ReplyToMessageId:    0,     // int32
		ReplyMarkup:         nil,   // ReplyMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SendDocument(t *testing.T) {
	rst, err := SendDocument(&botproto.SendDocument{
		ChatId:              "",    // string
		Document:            "",    // string
		Thumb:               "",    // string
		Caption:             "",    // string
		ParseMode:           "",    // string
		DisableNotification: false, // bool
		ReplyToMessageId:    0,     // int32
		ReplyMarkup:         nil,   // ReplyMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SendVideo(t *testing.T) {
	rst, err := SendVideo(&botproto.SendVideo{
		ChatId:              "",    // string
		Video:               "",    // string
		Duration:            0,     // int32
		Width:               0,     // int32
		Height:              0,     // int32
		Thumb:               "",    // string
		Caption:             "",    // string
		ParseMode:           "",    // string
		SupportsStreaming:   false, // bool
		DisableNotification: false, // bool
		ReplyToMessageId:    0,     // int32
		ReplyMarkup:         nil,   // ReplyMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SendAnimation(t *testing.T) {
	rst, err := SendAnimation(&botproto.SendAnimation{
		ChatId:              "",    // string
		Animation:           "",    // string
		Duration:            0,     // int32
		Width:               0,     // int32
		Height:              0,     // int32
		Thumb:               "",    // string
		Caption:             "",    // string
		ParseMode:           "",    // string
		DisableNotification: false, // bool
		ReplyToMessageId:    0,     // int32
		ReplyMarkup:         nil,   // ReplyMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SendVoice(t *testing.T) {
	rst, err := SendVoice(&botproto.SendVoice{
		ChatId:              "",    // string
		Voice:               "",    // string
		Caption:             "",    // string
		ParseMode:           "",    // string
		Duration:            0,     // int32
		DisableNotification: false, // bool
		ReplyToMessageId:    0,     // int32
		ReplyMarkup:         nil,   // ReplyMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SendVideoNote(t *testing.T) {
	rst, err := SendVideoNote(&botproto.SendVideoNote{
		ChatId:              "",    // string
		VideoNote:           "",    // string
		Duration:            0,     // int32
		Length:              0,     // int32
		Thumb:               "",    // string
		DisableNotification: false, // bool
		ReplyToMessageId:    0,     // int32
		ReplyMarkup:         nil,   // ReplyMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SendMediaGroup(t *testing.T) {
	rst, err := SendMediaGroup(&botproto.SendMediaGroup{
		ChatId:              "",    // string
		Media:               nil,   // repeated InputMedia
		DisableNotification: false, // bool
		ReplyToMessageId:    0,     // int32
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SendLocation(t *testing.T) {
	rst, err := SendLocation(&botproto.SendLocation{
		ChatId:              "",    // string
		Latitude:            0,     // int64
		Longitude:           0,     // int64
		LivePeriod:          0,     // int32
		DisableNotification: false, // bool
		ReplyToMessageId:    0,     // int32
		ReplyMarkup:         nil,   // ReplyMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_EditMessageLiveLocation(t *testing.T) {
	err := EditMessageLiveLocation(&botproto.EditMessageLiveLocation{
		ChatId:          "",  // string
		MessageId:       0,   // int32
		InlineMessageId: "",  // string
		Latitude:        0,   // int64
		Longitude:       0,   // int64
		ReplyMarkup:     nil, // InlineKeyboardMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
}

func TestBot_StopMessageLiveLocation(t *testing.T) {
	err := StopMessageLiveLocation(&botproto.StopMessageLiveLocation{
		ChatId:          "",  // string
		MessageId:       0,   // int32
		InlineMessageId: "",  // string
		ReplyMarkup:     nil, // InlineKeyboardMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
}

func TestBot_SendVenue(t *testing.T) {
	rst, err := SendVenue(&botproto.SendVenue{
		ChatId:              "",    // string
		Latitude:            0,     // int64
		Longitude:           0,     // int64
		Title:               "",    // string
		Address:             "",    // string
		FoursquareId:        "",    // string
		FoursquareType:      "",    // string
		DisableNotification: false, // bool
		ReplyToMessageId:    0,     // int32
		ReplyMarkup:         nil,   // ReplyMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SendContact(t *testing.T) {
	rst, err := SendContact(&botproto.SendContact{
		ChatId:              "",    // string
		PhoneNumber:         "",    // string
		FirstName:           "",    // string
		LastName:            "",    // string
		Vcard:               "",    // string
		DisableNotification: false, // bool
		ReplyToMessageId:    0,     // int32
		ReplyMarkup:         nil,   // ReplyMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SendPoll(t *testing.T) {
	rst, err := SendPoll(&botproto.SendPoll{
		ChatId:              "",    // string
		Question:            "",    // string
		Options:             nil,   // repeated string
		DisableNotification: false, // bool
		ReplyToMessageId:    0,     // int32
		ReplyMarkup:         nil,   // ReplyMarkup
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SendChatAction(t *testing.T) {
	rst, err := SendChatAction(&botproto.SendChatAction{
		ChatId: "", // string
		Action: "", // string
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_GetUserProfilePhotos(t *testing.T) {
	err := GetUserProfilePhotos(&botproto.GetUserProfilePhotos{
		UserId: 0, // int32
		Offset: 0, // int32
		Limit:  0, // int32
	})
	log.Println(assertions.ShouldBeNil(err))
}

func TestBot_GetFile(t *testing.T) {
	err := GetFile(&botproto.GetFile{
		FileId: "", // string
	})
	log.Println(assertions.ShouldBeNil(err))
}

func TestBot_KickChatMember(t *testing.T) {
	rst, err := KickChatMember(&botproto.KickChatMember{
		ChatId:    "", // string
		UserId:    0,  // int32
		UntilDate: 0,  // int32
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_UnbanChatMember(t *testing.T) {
	rst, err := UnbanChatMember(&botproto.UnbanChatMember{
		ChatId: "", // string
		UserId: 0,  // int32
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_RestrictChatMember(t *testing.T) {
	rst, err := RestrictChatMember(&botproto.RestrictChatMember{
		ChatId:      "",  // string
		UserId:      0,   // int32
		Permissions: nil, // ChatPermissions
		UntilDate:   0,   // int32
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_PromoteChatMember(t *testing.T) {
	rst, err := PromoteChatMember(&botproto.PromoteChatMember{
		ChatId:             "",    // string
		UserId:             0,     // int32
		CanChangeInfo:      false, // bool
		CanPostMessages:    false, // bool
		CanEditMessages:    false, // bool
		CanDeleteMessages:  false, // bool
		CanInviteUsers:     false, // bool
		CanRestrictMembers: false, // bool
		CanPinMessages:     false, // bool
		CanPromoteMembers:  false, // bool
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SetChatPermissions(t *testing.T) {
	rst, err := SetChatPermissions(&botproto.SetChatPermissions{
		ChatId:      "",  // string
		Permissions: nil, // ChatPermissions
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_ExportChatInviteLink(t *testing.T) {
	rst, err := ExportChatInviteLink(&botproto.ExportChatInviteLink{
		ChatId: "", // string
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SetChatPhoto(t *testing.T) {
	rst, err := SetChatPhoto(&botproto.SetChatPhoto{
		ChatId: "",  // string
		Photo:  nil, // InputFile
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_DeleteChatPhoto(t *testing.T) {
	rst, err := DeleteChatPhoto(&botproto.DeleteChatPhoto{
		ChatId: "", // string
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SetChatTitle(t *testing.T) {
	rst, err := SetChatTitle(&botproto.SetChatTitle{
		ChatId: "", // string
		Title:  "", // string
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SetChatDescription(t *testing.T) {
	rst, err := SetChatDescription(&botproto.SetChatDescription{
		ChatId:      "", // string
		Description: "", // string
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_PinChatMessage(t *testing.T) {
	rst, err := PinChatMessage(&botproto.PinChatMessage{
		ChatId:              "",    // string
		MessageId:           0,     // int32
		DisableNotification: false, // bool
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_UnpinChatMessage(t *testing.T) {
	rst, err := UnpinChatMessage(&botproto.UnpinChatMessage{
		ChatId: "", // string
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_LeaveChat(t *testing.T) {
	rst, err := LeaveChat(&botproto.LeaveChat{
		ChatId: "", // string
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_GetChat(t *testing.T) {
	rst, err := GetChat(&botproto.GetChat{
		ChatId: "", // string
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_GetChatAdministrators(t *testing.T) {
	rst, err := GetChatAdministrators(&botproto.GetChatAdministrators{
		ChatId: "", // string
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_GetChatMembersCount(t *testing.T) {
	rst, err := GetChatMembersCount(&botproto.GetChatMembersCount{
		ChatId: "", // string
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_GetChatMember(t *testing.T) {
	rst, err := GetChatMember(&botproto.GetChatMember{
		ChatId: "", // string
		UserId: 0,  // int32
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_SetChatStickerSet(t *testing.T) {
	rst, err := SetChatStickerSet(&botproto.SetChatStickerSet{
		ChatId:         "", // string
		StickerSetName: "", // string
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}

func TestBot_DeleteChatStickerSet(t *testing.T) {
	rst, err := DeleteChatStickerSet(&botproto.DeleteChatStickerSet{
		ChatId: "", // string
	})
	log.Println(assertions.ShouldBeNil(err))
	if nil == err {
		log.Printf("%#v", rst)
	}
}
