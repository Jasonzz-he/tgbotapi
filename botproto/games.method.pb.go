// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: games.method.proto

package botproto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// return Message
// Use this method to send a game. On success, the sent Message is returned.
type SendGame struct {
	ChatId              int32                 `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	GameShortName       string                `protobuf:"bytes,2,opt,name=game_short_name,json=gameShortName,proto3" json:"game_short_name,omitempty"`
	DisableNotification bool                  `protobuf:"varint,3,opt,name=disable_notification,json=disableNotification,proto3" json:"disable_notification,omitempty"`
	ReplyToMessageId    int32                 `protobuf:"varint,4,opt,name=reply_to_message_id,json=replyToMessageId,proto3" json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `protobuf:"bytes,5,opt,name=reply_markup,json=replyMarkup" json:"reply_markup,omitempty"`
}

func (m *SendGame) Reset()                    { *m = SendGame{} }
func (m *SendGame) String() string            { return proto.CompactTextString(m) }
func (*SendGame) ProtoMessage()               {}
func (*SendGame) Descriptor() ([]byte, []int) { return fileDescriptorGamesMethod, []int{0} }

func (m *SendGame) GetChatId() int32 {
	if m != nil {
		return m.ChatId
	}
	return 0
}

func (m *SendGame) GetGameShortName() string {
	if m != nil {
		return m.GameShortName
	}
	return ""
}

func (m *SendGame) GetDisableNotification() bool {
	if m != nil {
		return m.DisableNotification
	}
	return false
}

func (m *SendGame) GetReplyToMessageId() int32 {
	if m != nil {
		return m.ReplyToMessageId
	}
	return 0
}

func (m *SendGame) GetReplyMarkup() *InlineKeyboardMarkup {
	if m != nil {
		return m.ReplyMarkup
	}
	return nil
}

// return
// Use this method to set the score of the specified user in a game. On success, if the message was sent by the bot, returns the edited Message, otherwise returns True. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
type SetGameScore struct {
	UserId             int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Score              int32  `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	Force              bool   `protobuf:"varint,3,opt,name=force,proto3" json:"force,omitempty"`
	DisableEditMessage bool   `protobuf:"varint,4,opt,name=disable_edit_message,json=disableEditMessage,proto3" json:"disable_edit_message,omitempty"`
	ChatId             int32  `protobuf:"varint,5,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	MessageId          int32  `protobuf:"varint,6,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	InlineMessageId    string `protobuf:"bytes,7,opt,name=inline_message_id,json=inlineMessageId,proto3" json:"inline_message_id,omitempty"`
}

func (m *SetGameScore) Reset()                    { *m = SetGameScore{} }
func (m *SetGameScore) String() string            { return proto.CompactTextString(m) }
func (*SetGameScore) ProtoMessage()               {}
func (*SetGameScore) Descriptor() ([]byte, []int) { return fileDescriptorGamesMethod, []int{1} }

func (m *SetGameScore) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SetGameScore) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *SetGameScore) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

func (m *SetGameScore) GetDisableEditMessage() bool {
	if m != nil {
		return m.DisableEditMessage
	}
	return false
}

func (m *SetGameScore) GetChatId() int32 {
	if m != nil {
		return m.ChatId
	}
	return 0
}

func (m *SetGameScore) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *SetGameScore) GetInlineMessageId() string {
	if m != nil {
		return m.InlineMessageId
	}
	return ""
}

// return Array of GameHighScore
// Use this method to get data for high score tables. Will return the score of the specified user and several of his neighbors in a game. On success, returns an Array of GameHighScore objects.
// This method will currently return scores for the target user, plus two of his closest neighbors on each side. Will also return the top three users if the user and his neighbors are not among them. Please note that this behavior is subject to change.
//
type GetGameHighScores struct {
	UserId          int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChatId          int32  `protobuf:"varint,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	MessageId       int32  `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	InlineMessageId string `protobuf:"bytes,4,opt,name=inline_message_id,json=inlineMessageId,proto3" json:"inline_message_id,omitempty"`
}

func (m *GetGameHighScores) Reset()                    { *m = GetGameHighScores{} }
func (m *GetGameHighScores) String() string            { return proto.CompactTextString(m) }
func (*GetGameHighScores) ProtoMessage()               {}
func (*GetGameHighScores) Descriptor() ([]byte, []int) { return fileDescriptorGamesMethod, []int{2} }

func (m *GetGameHighScores) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetGameHighScores) GetChatId() int32 {
	if m != nil {
		return m.ChatId
	}
	return 0
}

func (m *GetGameHighScores) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *GetGameHighScores) GetInlineMessageId() string {
	if m != nil {
		return m.InlineMessageId
	}
	return ""
}

func init() {
	proto.RegisterType((*SendGame)(nil), "botproto.sendGame")
	proto.RegisterType((*SetGameScore)(nil), "botproto.setGameScore")
	proto.RegisterType((*GetGameHighScores)(nil), "botproto.getGameHighScores")
}
func (m *SendGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendGame) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ChatId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGamesMethod(dAtA, i, uint64(m.ChatId))
	}
	if len(m.GameShortName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGamesMethod(dAtA, i, uint64(len(m.GameShortName)))
		i += copy(dAtA[i:], m.GameShortName)
	}
	if m.DisableNotification {
		dAtA[i] = 0x18
		i++
		if m.DisableNotification {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ReplyToMessageId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintGamesMethod(dAtA, i, uint64(m.ReplyToMessageId))
	}
	if m.ReplyMarkup != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintGamesMethod(dAtA, i, uint64(m.ReplyMarkup.Size()))
		n1, err := m.ReplyMarkup.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *SetGameScore) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetGameScore) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UserId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGamesMethod(dAtA, i, uint64(m.UserId))
	}
	if m.Score != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGamesMethod(dAtA, i, uint64(m.Score))
	}
	if m.Force {
		dAtA[i] = 0x18
		i++
		if m.Force {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.DisableEditMessage {
		dAtA[i] = 0x20
		i++
		if m.DisableEditMessage {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ChatId != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintGamesMethod(dAtA, i, uint64(m.ChatId))
	}
	if m.MessageId != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintGamesMethod(dAtA, i, uint64(m.MessageId))
	}
	if len(m.InlineMessageId) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintGamesMethod(dAtA, i, uint64(len(m.InlineMessageId)))
		i += copy(dAtA[i:], m.InlineMessageId)
	}
	return i, nil
}

func (m *GetGameHighScores) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetGameHighScores) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UserId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGamesMethod(dAtA, i, uint64(m.UserId))
	}
	if m.ChatId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGamesMethod(dAtA, i, uint64(m.ChatId))
	}
	if m.MessageId != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGamesMethod(dAtA, i, uint64(m.MessageId))
	}
	if len(m.InlineMessageId) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGamesMethod(dAtA, i, uint64(len(m.InlineMessageId)))
		i += copy(dAtA[i:], m.InlineMessageId)
	}
	return i, nil
}

func encodeFixed64GamesMethod(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32GamesMethod(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGamesMethod(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SendGame) Size() (n int) {
	var l int
	_ = l
	if m.ChatId != 0 {
		n += 1 + sovGamesMethod(uint64(m.ChatId))
	}
	l = len(m.GameShortName)
	if l > 0 {
		n += 1 + l + sovGamesMethod(uint64(l))
	}
	if m.DisableNotification {
		n += 2
	}
	if m.ReplyToMessageId != 0 {
		n += 1 + sovGamesMethod(uint64(m.ReplyToMessageId))
	}
	if m.ReplyMarkup != nil {
		l = m.ReplyMarkup.Size()
		n += 1 + l + sovGamesMethod(uint64(l))
	}
	return n
}

func (m *SetGameScore) Size() (n int) {
	var l int
	_ = l
	if m.UserId != 0 {
		n += 1 + sovGamesMethod(uint64(m.UserId))
	}
	if m.Score != 0 {
		n += 1 + sovGamesMethod(uint64(m.Score))
	}
	if m.Force {
		n += 2
	}
	if m.DisableEditMessage {
		n += 2
	}
	if m.ChatId != 0 {
		n += 1 + sovGamesMethod(uint64(m.ChatId))
	}
	if m.MessageId != 0 {
		n += 1 + sovGamesMethod(uint64(m.MessageId))
	}
	l = len(m.InlineMessageId)
	if l > 0 {
		n += 1 + l + sovGamesMethod(uint64(l))
	}
	return n
}

func (m *GetGameHighScores) Size() (n int) {
	var l int
	_ = l
	if m.UserId != 0 {
		n += 1 + sovGamesMethod(uint64(m.UserId))
	}
	if m.ChatId != 0 {
		n += 1 + sovGamesMethod(uint64(m.ChatId))
	}
	if m.MessageId != 0 {
		n += 1 + sovGamesMethod(uint64(m.MessageId))
	}
	l = len(m.InlineMessageId)
	if l > 0 {
		n += 1 + l + sovGamesMethod(uint64(l))
	}
	return n
}

func sovGamesMethod(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGamesMethod(x uint64) (n int) {
	return sovGamesMethod(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGamesMethod
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: sendGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: sendGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChatId", wireType)
			}
			m.ChatId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChatId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameShortName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGamesMethod
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GameShortName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableNotification", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DisableNotification = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplyToMessageId", wireType)
			}
			m.ReplyToMessageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReplyToMessageId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplyMarkup", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGamesMethod
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ReplyMarkup == nil {
				m.ReplyMarkup = &InlineKeyboardMarkup{}
			}
			if err := m.ReplyMarkup.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGamesMethod(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGamesMethod
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SetGameScore) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGamesMethod
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: setGameScore: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: setGameScore: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Score |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Force", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Force = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableEditMessage", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DisableEditMessage = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChatId", wireType)
			}
			m.ChatId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChatId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageId", wireType)
			}
			m.MessageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InlineMessageId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGamesMethod
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InlineMessageId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGamesMethod(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGamesMethod
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetGameHighScores) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGamesMethod
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: getGameHighScores: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: getGameHighScores: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChatId", wireType)
			}
			m.ChatId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChatId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageId", wireType)
			}
			m.MessageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InlineMessageId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGamesMethod
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InlineMessageId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGamesMethod(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGamesMethod
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGamesMethod(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGamesMethod
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGamesMethod
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGamesMethod
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGamesMethod
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGamesMethod(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGamesMethod = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGamesMethod   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("games.method.proto", fileDescriptorGamesMethod) }

var fileDescriptorGamesMethod = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x8a, 0xdb, 0x30,
	0x18, 0xc5, 0x51, 0x12, 0x27, 0x8e, 0x92, 0x36, 0x8d, 0x12, 0xa8, 0x29, 0xd4, 0x98, 0x2c, 0x8a,
	0x29, 0xd4, 0xf4, 0xcf, 0x09, 0x5a, 0x28, 0xad, 0x29, 0xc9, 0xc2, 0xe9, 0xde, 0xc8, 0x96, 0x62,
	0x8b, 0xc6, 0x96, 0x91, 0x94, 0x45, 0x6e, 0xd1, 0x63, 0x75, 0xd9, 0x23, 0x0c, 0xd9, 0xcd, 0x6e,
	0x8e, 0x30, 0x48, 0x72, 0x32, 0x1e, 0x98, 0x19, 0x66, 0xe7, 0xf7, 0xbd, 0xcf, 0xd2, 0xfb, 0x3d,
	0x41, 0x54, 0xe0, 0x8a, 0xca, 0xa8, 0xa2, 0xaa, 0xe4, 0x24, 0x6a, 0x04, 0x57, 0x1c, 0xb9, 0x19,
	0x57, 0xe6, 0xeb, 0xcd, 0x4b, 0x55, 0x64, 0x5c, 0xe1, 0x86, 0x59, 0x67, 0x75, 0x03, 0xa0, 0x2b,
	0x69, 0x4d, 0x7e, 0xe0, 0x8a, 0xa2, 0xd7, 0x70, 0x94, 0x97, 0x58, 0xa5, 0x8c, 0x78, 0x20, 0x00,
	0xa1, 0x93, 0x0c, 0xb5, 0x8c, 0x09, 0x7a, 0x07, 0x67, 0xfa, 0xd4, 0x54, 0x96, 0x5c, 0xa8, 0xb4,
	0xc6, 0x15, 0xf5, 0x7a, 0x01, 0x08, 0xc7, 0xc9, 0x0b, 0x3d, 0xde, 0xea, 0xe9, 0x46, 0x1f, 0xf0,
	0x09, 0x2e, 0x09, 0x93, 0x38, 0xdb, 0xd3, 0xb4, 0xe6, 0x8a, 0xed, 0x58, 0x8e, 0x15, 0xe3, 0xb5,
	0xd7, 0x0f, 0x40, 0xe8, 0x26, 0x8b, 0xd6, 0xdb, 0x74, 0x2c, 0xf4, 0x01, 0x2e, 0x04, 0x6d, 0xf6,
	0xc7, 0x54, 0xf1, 0xb4, 0xa2, 0x52, 0xe2, 0x82, 0xea, 0xfb, 0x07, 0xe6, 0xfe, 0x57, 0xc6, 0xfa,
	0xcd, 0xd7, 0xd6, 0x88, 0x09, 0xfa, 0x0a, 0xa7, 0x76, 0xbd, 0xc2, 0xe2, 0xcf, 0xa1, 0xf1, 0x9c,
	0x00, 0x84, 0x93, 0xcf, 0x7e, 0x74, 0x06, 0x8c, 0xe2, 0x7a, 0xcf, 0x6a, 0xfa, 0x8b, 0x1e, 0x33,
	0x8e, 0x05, 0x59, 0x9b, 0xad, 0x64, 0x62, 0xfe, 0xb1, 0x62, 0x75, 0x0d, 0xe0, 0x54, 0x52, 0xa5,
	0x89, 0xb7, 0x39, 0x17, 0x06, 0xfb, 0x20, 0xa9, 0xe8, 0x60, 0x6b, 0x19, 0x13, 0xb4, 0x84, 0x8e,
	0xd4, 0x1b, 0x06, 0xd6, 0x49, 0xac, 0xd0, 0xd3, 0x1d, 0x17, 0x39, 0x6d, 0xa9, 0xac, 0x40, 0x1f,
	0xef, 0xd0, 0x29, 0x61, 0xea, 0xcc, 0x62, 0x40, 0xdc, 0x04, 0xb5, 0xde, 0x77, 0xc2, 0x54, 0x0b,
	0xd3, 0x6d, 0xdb, 0xb9, 0xd7, 0xf6, 0x5b, 0x08, 0x3b, 0x4d, 0x0c, 0x8d, 0x37, 0xae, 0x2e, 0x15,
	0xbc, 0x87, 0x73, 0x66, 0x20, 0xbb, 0x7d, 0x8d, 0xcc, 0x73, 0xcc, 0xac, 0x71, 0xa9, 0x6b, 0xf5,
	0x17, 0xc0, 0x79, 0x61, 0x59, 0x7f, 0xb2, 0xa2, 0x34, 0xbc, 0xf2, 0x71, 0xe0, 0x4e, 0xa4, 0xde,
	0x13, 0x91, 0xfa, 0xcf, 0x8a, 0x34, 0x78, 0x30, 0xd2, 0xb7, 0xe9, 0xbf, 0x93, 0x0f, 0xfe, 0x9f,
	0x7c, 0x70, 0x75, 0xf2, 0x41, 0x36, 0x34, 0xaf, 0xf6, 0xe5, 0x36, 0x00, 0x00, 0xff, 0xff, 0x60,
	0xc8, 0xee, 0x94, 0xb6, 0x02, 0x00, 0x00,
}