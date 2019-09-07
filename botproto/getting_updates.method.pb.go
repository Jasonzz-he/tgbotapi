// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: getting_updates.method.proto

package botproto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// return
// Use this method to receive incoming updates using long polling (wiki). An Array of Update objects is returned.
// Notes1. This method will not work if an outgoing webhook is set up.2. In order to avoid getting duplicate updates, recalculate offset after each server response.
//
type GetUpdates struct {
	Offset         int32    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit          int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Timeout        int32    `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	AllowedUpdates []string `protobuf:"bytes,4,rep,name=allowed_updates,json=allowedUpdates" json:"allowed_updates,omitempty"`
}

func (m *GetUpdates) Reset()                    { *m = GetUpdates{} }
func (m *GetUpdates) String() string            { return proto.CompactTextString(m) }
func (*GetUpdates) ProtoMessage()               {}
func (*GetUpdates) Descriptor() ([]byte, []int) { return fileDescriptorGettingUpdatesMethod, []int{0} }

func (m *GetUpdates) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetUpdates) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetUpdates) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *GetUpdates) GetAllowedUpdates() []string {
	if m != nil {
		return m.AllowedUpdates
	}
	return nil
}

// return True
// Use this method to specify a url and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
// If you'd like to make sure that the Webhook request comes from Telegram, we recommend using a secret path in the URL, e.g. https://www.example.com/<token>. Since nobody else knows your bot‘s token, you can be pretty sure it’s us.
// Notes1. You will not be able to receive updates using getUpdates for as long as an outgoing webhook is set up.2. To use a self-signed certificate, you need to upload your public key certificate using certificate parameter. Please upload as InputFile, sending a String will not work.3. Ports currently supported for Webhooks: 443, 80, 88, 8443.
// NEW! If you're having any trouble setting up webhooks, please check out this amazing guide to Webhooks.
//
type SetWebhook struct {
	Url            string     `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Certificate    *InputFile `protobuf:"bytes,2,opt,name=certificate" json:"certificate,omitempty"`
	MaxConnections int32      `protobuf:"varint,3,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	AllowedUpdates []string   `protobuf:"bytes,4,rep,name=allowed_updates,json=allowedUpdates" json:"allowed_updates,omitempty"`
}

func (m *SetWebhook) Reset()                    { *m = SetWebhook{} }
func (m *SetWebhook) String() string            { return proto.CompactTextString(m) }
func (*SetWebhook) ProtoMessage()               {}
func (*SetWebhook) Descriptor() ([]byte, []int) { return fileDescriptorGettingUpdatesMethod, []int{1} }

func (m *SetWebhook) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SetWebhook) GetCertificate() *InputFile {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *SetWebhook) GetMaxConnections() int32 {
	if m != nil {
		return m.MaxConnections
	}
	return 0
}

func (m *SetWebhook) GetAllowedUpdates() []string {
	if m != nil {
		return m.AllowedUpdates
	}
	return nil
}

// return True
// Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success. Requires no parameters.
type DeleteWebhook struct {
}

func (m *DeleteWebhook) Reset()         { *m = DeleteWebhook{} }
func (m *DeleteWebhook) String() string { return proto.CompactTextString(m) }
func (*DeleteWebhook) ProtoMessage()    {}
func (*DeleteWebhook) Descriptor() ([]byte, []int) {
	return fileDescriptorGettingUpdatesMethod, []int{2}
}

// return
// Use this method to get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
type GetWebhookInfo struct {
}

func (m *GetWebhookInfo) Reset()         { *m = GetWebhookInfo{} }
func (m *GetWebhookInfo) String() string { return proto.CompactTextString(m) }
func (*GetWebhookInfo) ProtoMessage()    {}
func (*GetWebhookInfo) Descriptor() ([]byte, []int) {
	return fileDescriptorGettingUpdatesMethod, []int{3}
}

func init() {
	proto.RegisterType((*GetUpdates)(nil), "botproto.getUpdates")
	proto.RegisterType((*SetWebhook)(nil), "botproto.setWebhook")
	proto.RegisterType((*DeleteWebhook)(nil), "botproto.deleteWebhook")
	proto.RegisterType((*GetWebhookInfo)(nil), "botproto.getWebhookInfo")
}
func (m *GetUpdates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetUpdates) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Offset != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGettingUpdatesMethod(dAtA, i, uint64(m.Offset))
	}
	if m.Limit != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGettingUpdatesMethod(dAtA, i, uint64(m.Limit))
	}
	if m.Timeout != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGettingUpdatesMethod(dAtA, i, uint64(m.Timeout))
	}
	if len(m.AllowedUpdates) > 0 {
		for _, s := range m.AllowedUpdates {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *SetWebhook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetWebhook) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGettingUpdatesMethod(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if m.Certificate != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGettingUpdatesMethod(dAtA, i, uint64(m.Certificate.Size()))
		n1, err := m.Certificate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.MaxConnections != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGettingUpdatesMethod(dAtA, i, uint64(m.MaxConnections))
	}
	if len(m.AllowedUpdates) > 0 {
		for _, s := range m.AllowedUpdates {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *DeleteWebhook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteWebhook) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetWebhookInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetWebhookInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeFixed64GettingUpdatesMethod(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32GettingUpdatesMethod(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGettingUpdatesMethod(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GetUpdates) Size() (n int) {
	var l int
	_ = l
	if m.Offset != 0 {
		n += 1 + sovGettingUpdatesMethod(uint64(m.Offset))
	}
	if m.Limit != 0 {
		n += 1 + sovGettingUpdatesMethod(uint64(m.Limit))
	}
	if m.Timeout != 0 {
		n += 1 + sovGettingUpdatesMethod(uint64(m.Timeout))
	}
	if len(m.AllowedUpdates) > 0 {
		for _, s := range m.AllowedUpdates {
			l = len(s)
			n += 1 + l + sovGettingUpdatesMethod(uint64(l))
		}
	}
	return n
}

func (m *SetWebhook) Size() (n int) {
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovGettingUpdatesMethod(uint64(l))
	}
	if m.Certificate != nil {
		l = m.Certificate.Size()
		n += 1 + l + sovGettingUpdatesMethod(uint64(l))
	}
	if m.MaxConnections != 0 {
		n += 1 + sovGettingUpdatesMethod(uint64(m.MaxConnections))
	}
	if len(m.AllowedUpdates) > 0 {
		for _, s := range m.AllowedUpdates {
			l = len(s)
			n += 1 + l + sovGettingUpdatesMethod(uint64(l))
		}
	}
	return n
}

func (m *DeleteWebhook) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *GetWebhookInfo) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovGettingUpdatesMethod(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGettingUpdatesMethod(x uint64) (n int) {
	return sovGettingUpdatesMethod(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetUpdates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGettingUpdatesMethod
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
			return fmt.Errorf("proto: getUpdates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: getUpdates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGettingUpdatesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGettingUpdatesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGettingUpdatesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedUpdates", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGettingUpdatesMethod
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
				return ErrInvalidLengthGettingUpdatesMethod
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedUpdates = append(m.AllowedUpdates, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGettingUpdatesMethod(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGettingUpdatesMethod
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
func (m *SetWebhook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGettingUpdatesMethod
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
			return fmt.Errorf("proto: setWebhook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: setWebhook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGettingUpdatesMethod
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
				return ErrInvalidLengthGettingUpdatesMethod
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGettingUpdatesMethod
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
				return ErrInvalidLengthGettingUpdatesMethod
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Certificate == nil {
				m.Certificate = &InputFile{}
			}
			if err := m.Certificate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxConnections", wireType)
			}
			m.MaxConnections = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGettingUpdatesMethod
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxConnections |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedUpdates", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGettingUpdatesMethod
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
				return ErrInvalidLengthGettingUpdatesMethod
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedUpdates = append(m.AllowedUpdates, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGettingUpdatesMethod(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGettingUpdatesMethod
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
func (m *DeleteWebhook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGettingUpdatesMethod
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
			return fmt.Errorf("proto: deleteWebhook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: deleteWebhook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGettingUpdatesMethod(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGettingUpdatesMethod
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
func (m *GetWebhookInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGettingUpdatesMethod
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
			return fmt.Errorf("proto: getWebhookInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: getWebhookInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipGettingUpdatesMethod(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGettingUpdatesMethod
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
func skipGettingUpdatesMethod(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGettingUpdatesMethod
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
					return 0, ErrIntOverflowGettingUpdatesMethod
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
					return 0, ErrIntOverflowGettingUpdatesMethod
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
				return 0, ErrInvalidLengthGettingUpdatesMethod
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGettingUpdatesMethod
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
				next, err := skipGettingUpdatesMethod(dAtA[start:])
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
	ErrInvalidLengthGettingUpdatesMethod = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGettingUpdatesMethod   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("getting_updates.method.proto", fileDescriptorGettingUpdatesMethod) }

var fileDescriptorGettingUpdatesMethod = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x19, 0x6b, 0xab, 0x7d, 0xd5, 0xb4, 0x8c, 0x22, 0x41, 0xa4, 0x94, 0x6e, 0xec, 0x2a,
	0x0b, 0xc5, 0x0b, 0x28, 0x08, 0xdd, 0x06, 0xc4, 0x65, 0x99, 0x24, 0x2f, 0xe9, 0xe0, 0x64, 0x5e,
	0x48, 0x5e, 0xb0, 0x1b, 0xcf, 0xe3, 0x55, 0x5c, 0x7a, 0x04, 0xc9, 0x49, 0xa4, 0x93, 0x44, 0x5d,
	0xba, 0xfb, 0xff, 0x6f, 0x66, 0xf8, 0xbf, 0x81, 0xab, 0x0c, 0x99, 0xb5, 0xcd, 0x36, 0x75, 0x91,
	0x28, 0xc6, 0x2a, 0xc8, 0x91, 0xb7, 0x94, 0x04, 0x45, 0x49, 0x4c, 0xf2, 0x38, 0x22, 0x76, 0xe9,
	0xd2, 0xe3, 0x2c, 0x22, 0x56, 0x85, 0x6e, 0x4f, 0x96, 0x6f, 0x00, 0x19, 0xf2, 0x53, 0xfb, 0x48,
	0x5e, 0xc0, 0x88, 0xd2, 0xb4, 0x42, 0xf6, 0xc5, 0x42, 0xac, 0x86, 0x61, 0xd7, 0xe4, 0x39, 0x0c,
	0x8d, 0xce, 0x35, 0xfb, 0x07, 0x0e, 0xb7, 0x45, 0xfa, 0x70, 0xc4, 0x3a, 0x47, 0xaa, 0xd9, 0x1f,
	0x38, 0xde, 0x57, 0x79, 0x0d, 0x53, 0x65, 0x0c, 0xbd, 0x62, 0xd2, 0xfb, 0xf8, 0x87, 0x8b, 0xc1,
	0x6a, 0x1c, 0x7a, 0x1d, 0xee, 0x06, 0x97, 0xef, 0x02, 0xa0, 0x42, 0x7e, 0xc6, 0x68, 0x4b, 0xf4,
	0x22, 0x67, 0x30, 0xa8, 0x4b, 0xe3, 0xc6, 0xc7, 0xe1, 0x3e, 0xca, 0x3b, 0x98, 0xc4, 0x58, 0xb2,
	0x4e, 0x75, 0xac, 0x18, 0xdd, 0xfe, 0xe4, 0xe6, 0x2c, 0xe8, 0xff, 0x13, 0xac, 0x6d, 0x51, 0xf3,
	0xa3, 0x36, 0x18, 0xfe, 0xbd, 0xb7, 0x17, 0xc8, 0xd5, 0x6e, 0x13, 0x93, 0xb5, 0x18, 0xb3, 0x26,
	0x5b, 0x75, 0x8a, 0x5e, 0xae, 0x76, 0x0f, 0xbf, 0xf4, 0xff, 0xa6, 0x53, 0x38, 0x4d, 0xd0, 0x20,
	0x63, 0xe7, 0xba, 0x9c, 0x81, 0x97, 0xfd, 0x98, 0xaf, 0x6d, 0x4a, 0xf7, 0x27, 0x1f, 0xcd, 0x5c,
	0x7c, 0x36, 0x73, 0xf1, 0xd5, 0xcc, 0x45, 0x34, 0x72, 0x82, 0xb7, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xb1, 0x0e, 0xa8, 0x13, 0x9a, 0x01, 0x00, 0x00,
}
