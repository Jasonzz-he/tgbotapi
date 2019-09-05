package tgbotapi

import "github.com/Jasonzz-he/tgbotapi/botproto"

func (b *Bot) GetMe() (*botproto.User, error) {
	var rst = new(botproto.User)
	err := b.Post("getMe", nil, rst)
	return rst, err
}

func GetMe() (*botproto.User, error) {
	return globalBot.GetMe()
}

func (b *Bot) SendMessage(in *botproto.SendMessage) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("sendMessage", in, rst)
	return rst, err
}

func SendMessage(in *botproto.SendMessage) (*botproto.Message, error) {
	return globalBot.SendMessage(in)
}

func (b *Bot) ForwardMessage(in *botproto.ForwardMessage) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("forwardMessage", in, rst)
	return rst, err
}

func ForwardMessage(in *botproto.ForwardMessage) (*botproto.Message, error) {
	return globalBot.ForwardMessage(in)
}

func (b *Bot) SendPhoto(in *botproto.SendPhoto) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("sendPhoto", in, rst)
	return rst, err
}

func SendPhoto(in *botproto.SendPhoto) (*botproto.Message, error) {
	return globalBot.SendPhoto(in)
}

func (b *Bot) SendAudio(in *botproto.SendAudio) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("sendAudio", in, rst)
	return rst, err
}

func SendAudio(in *botproto.SendAudio) (*botproto.Message, error) {
	return globalBot.SendAudio(in)
}

func (b *Bot) SendDocument(in *botproto.SendDocument) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("sendDocument", in, rst)
	return rst, err
}

func SendDocument(in *botproto.SendDocument) (*botproto.Message, error) {
	return globalBot.SendDocument(in)
}

func (b *Bot) SendVideo(in *botproto.SendVideo) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("sendVideo", in, rst)
	return rst, err
}

func SendVideo(in *botproto.SendVideo) (*botproto.Message, error) {
	return globalBot.SendVideo(in)
}

func (b *Bot) SendAnimation(in *botproto.SendAnimation) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("sendAnimation", in, rst)
	return rst, err
}

func SendAnimation(in *botproto.SendAnimation) (*botproto.Message, error) {
	return globalBot.SendAnimation(in)
}

func (b *Bot) SendVoice(in *botproto.SendVoice) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("sendVoice", in, rst)
	return rst, err
}

func SendVoice(in *botproto.SendVoice) (*botproto.Message, error) {
	return globalBot.SendVoice(in)
}

func (b *Bot) SendVideoNote(in *botproto.SendVideoNote) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("sendVideoNote", in, rst)
	return rst, err
}

func SendVideoNote(in *botproto.SendVideoNote) (*botproto.Message, error) {
	return globalBot.SendVideoNote(in)
}

func (b *Bot) SendMediaGroup(in *botproto.SendMediaGroup) ([]*botproto.Message, error) {
	var rst = make([]*botproto.Message, 0)
	err := b.Post("sendMediaGroup", in, rst)
	return rst, err
}

func SendMediaGroup(in *botproto.SendMediaGroup) ([]*botproto.Message, error) {
	return globalBot.SendMediaGroup(in)
}

func (b *Bot) SendLocation(in *botproto.SendLocation) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("sendLocation", in, rst)
	return rst, err
}

func SendLocation(in *botproto.SendLocation) (*botproto.Message, error) {
	return globalBot.SendLocation(in)
}

func (b *Bot) EditMessageLiveLocation(in *botproto.EditMessageLiveLocation) error {
	err := b.Post("editMessageLiveLocation", in, nil)
	return err
}

func EditMessageLiveLocation(in *botproto.EditMessageLiveLocation) error {
	return globalBot.EditMessageLiveLocation(in)
}

func (b *Bot) StopMessageLiveLocation(in *botproto.StopMessageLiveLocation) error {
	err := b.Post("stopMessageLiveLocation", in, nil)
	return err
}

func StopMessageLiveLocation(in *botproto.StopMessageLiveLocation) error {
	return globalBot.StopMessageLiveLocation(in)
}

func (b *Bot) SendVenue(in *botproto.SendVenue) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("sendVenue", in, rst)
	return rst, err
}

func SendVenue(in *botproto.SendVenue) (*botproto.Message, error) {
	return globalBot.SendVenue(in)
}

func (b *Bot) SendContact(in *botproto.SendContact) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("sendContact", in, rst)
	return rst, err
}

func SendContact(in *botproto.SendContact) (*botproto.Message, error) {
	return globalBot.SendContact(in)
}

func (b *Bot) SendPoll(in *botproto.SendPoll) (*botproto.Message, error) {
	var rst = new(botproto.Message)
	err := b.Post("sendPoll", in, rst)
	return rst, err
}

func SendPoll(in *botproto.SendPoll) (*botproto.Message, error) {
	return globalBot.SendPoll(in)
}

func (b *Bot) SendChatAction(in *botproto.SendChatAction) (*bool, error) {
	var rst = new(bool)
	err := b.Post("sendChatAction", in, rst)
	return rst, err
}

func SendChatAction(in *botproto.SendChatAction) (*bool, error) {
	return globalBot.SendChatAction(in)
}

func (b *Bot) GetUserProfilePhotos(in *botproto.GetUserProfilePhotos) error {
	err := b.Post("getUserProfilePhotos", in, nil)
	return err
}

func GetUserProfilePhotos(in *botproto.GetUserProfilePhotos) error {
	return globalBot.GetUserProfilePhotos(in)
}

func (b *Bot) GetFile(in *botproto.GetFile) error {
	err := b.Post("getFile", in, nil)
	return err
}

func GetFile(in *botproto.GetFile) error {
	return globalBot.GetFile(in)
}

func (b *Bot) KickChatMember(in *botproto.KickChatMember) (*bool, error) {
	var rst = new(bool)
	err := b.Post("kickChatMember", in, rst)
	return rst, err
}

func KickChatMember(in *botproto.KickChatMember) (*bool, error) {
	return globalBot.KickChatMember(in)
}

func (b *Bot) UnbanChatMember(in *botproto.UnbanChatMember) (*bool, error) {
	var rst = new(bool)
	err := b.Post("unbanChatMember", in, rst)
	return rst, err
}

func UnbanChatMember(in *botproto.UnbanChatMember) (*bool, error) {
	return globalBot.UnbanChatMember(in)
}

func (b *Bot) RestrictChatMember(in *botproto.RestrictChatMember) (*bool, error) {
	var rst = new(bool)
	err := b.Post("restrictChatMember", in, rst)
	return rst, err
}

func RestrictChatMember(in *botproto.RestrictChatMember) (*bool, error) {
	return globalBot.RestrictChatMember(in)
}

func (b *Bot) PromoteChatMember(in *botproto.PromoteChatMember) (*bool, error) {
	var rst = new(bool)
	err := b.Post("promoteChatMember", in, rst)
	return rst, err
}

func PromoteChatMember(in *botproto.PromoteChatMember) (*bool, error) {
	return globalBot.PromoteChatMember(in)
}

func (b *Bot) SetChatPermissions(in *botproto.SetChatPermissions) (*bool, error) {
	var rst = new(bool)
	err := b.Post("setChatPermissions", in, rst)
	return rst, err
}

func SetChatPermissions(in *botproto.SetChatPermissions) (*bool, error) {
	return globalBot.SetChatPermissions(in)
}

func (b *Bot) ExportChatInviteLink(in *botproto.ExportChatInviteLink) (*string, error) {
	var rst = new(string)
	err := b.Post("exportChatInviteLink", in, rst)
	return rst, err
}

func ExportChatInviteLink(in *botproto.ExportChatInviteLink) (*string, error) {
	return globalBot.ExportChatInviteLink(in)
}

func (b *Bot) SetChatPhoto(in *botproto.SetChatPhoto) (*bool, error) {
	var rst = new(bool)
	err := b.Post("setChatPhoto", in, rst)
	return rst, err
}

func SetChatPhoto(in *botproto.SetChatPhoto) (*bool, error) {
	return globalBot.SetChatPhoto(in)
}

func (b *Bot) DeleteChatPhoto(in *botproto.DeleteChatPhoto) (*bool, error) {
	var rst = new(bool)
	err := b.Post("deleteChatPhoto", in, rst)
	return rst, err
}

func DeleteChatPhoto(in *botproto.DeleteChatPhoto) (*bool, error) {
	return globalBot.DeleteChatPhoto(in)
}

func (b *Bot) SetChatTitle(in *botproto.SetChatTitle) (*bool, error) {
	var rst = new(bool)
	err := b.Post("setChatTitle", in, rst)
	return rst, err
}

func SetChatTitle(in *botproto.SetChatTitle) (*bool, error) {
	return globalBot.SetChatTitle(in)
}

func (b *Bot) SetChatDescription(in *botproto.SetChatDescription) (*bool, error) {
	var rst = new(bool)
	err := b.Post("setChatDescription", in, rst)
	return rst, err
}

func SetChatDescription(in *botproto.SetChatDescription) (*bool, error) {
	return globalBot.SetChatDescription(in)
}

func (b *Bot) PinChatMessage(in *botproto.PinChatMessage) (*bool, error) {
	var rst = new(bool)
	err := b.Post("pinChatMessage", in, rst)
	return rst, err
}

func PinChatMessage(in *botproto.PinChatMessage) (*bool, error) {
	return globalBot.PinChatMessage(in)
}

func (b *Bot) UnpinChatMessage(in *botproto.UnpinChatMessage) (*bool, error) {
	var rst = new(bool)
	err := b.Post("unpinChatMessage", in, rst)
	return rst, err
}

func UnpinChatMessage(in *botproto.UnpinChatMessage) (*bool, error) {
	return globalBot.UnpinChatMessage(in)
}

func (b *Bot) LeaveChat(in *botproto.LeaveChat) (*bool, error) {
	var rst = new(bool)
	err := b.Post("leaveChat", in, rst)
	return rst, err
}

func LeaveChat(in *botproto.LeaveChat) (*bool, error) {
	return globalBot.LeaveChat(in)
}

func (b *Bot) GetChat(in *botproto.GetChat) (*botproto.Chat, error) {
	var rst = new(botproto.Chat)
	err := b.Post("getChat", in, rst)
	return rst, err
}

func GetChat(in *botproto.GetChat) (*botproto.Chat, error) {
	return globalBot.GetChat(in)
}

func (b *Bot) GetChatAdministrators(in *botproto.GetChatAdministrators) ([]*botproto.ChatMember, error) {
	var rst = make([]*botproto.ChatMember, 0)
	err := b.Post("getChatAdministrators", in, rst)
	return rst, err
}

func GetChatAdministrators(in *botproto.GetChatAdministrators) ([]*botproto.ChatMember, error) {
	return globalBot.GetChatAdministrators(in)
}

func (b *Bot) GetChatMembersCount(in *botproto.GetChatMembersCount) (*int32, error) {
	var rst = new(int32)
	err := b.Post("getChatMembersCount", in, rst)
	return rst, err
}

func GetChatMembersCount(in *botproto.GetChatMembersCount) (*int32, error) {
	return globalBot.GetChatMembersCount(in)
}

func (b *Bot) GetChatMember(in *botproto.GetChatMember) (*botproto.ChatMember, error) {
	var rst = new(botproto.ChatMember)
	err := b.Post("getChatMember", in, rst)
	return rst, err
}

func GetChatMember(in *botproto.GetChatMember) (*botproto.ChatMember, error) {
	return globalBot.GetChatMember(in)
}

func (b *Bot) SetChatStickerSet(in *botproto.SetChatStickerSet) (*bool, error) {
	var rst = new(bool)
	err := b.Post("setChatStickerSet", in, rst)
	return rst, err
}

func SetChatStickerSet(in *botproto.SetChatStickerSet) (*bool, error) {
	return globalBot.SetChatStickerSet(in)
}

func (b *Bot) DeleteChatStickerSet(in *botproto.DeleteChatStickerSet) (*bool, error) {
	var rst = new(bool)
	err := b.Post("deleteChatStickerSet", in, rst)
	return rst, err
}

func DeleteChatStickerSet(in *botproto.DeleteChatStickerSet) (*bool, error) {
	return globalBot.DeleteChatStickerSet(in)
}
