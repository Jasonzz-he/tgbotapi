package main

import (
	"log"
	"strings"

	"github.com/Jasonzz-he/golanglib/generate"
)

/*
   创建日期：2019/9/4
   修改日期：2019/9/4
   作者：Jason
*/

var ps = make([]*ProtoFile, 0)

type ProtoFile struct {
	ModelName        string
	ProtoMessageList []*ProtoMessage
	ProtoRPCList     []*ProtoRPC
	curMessage       *ProtoMessage
	curRPC           *ProtoRPC
}

func NewProtoFile(modelName string) *ProtoFile {
	p := &ProtoFile{
		ModelName: modelName,
	}
	ps = append(ps, p)
	return p
}

func (p *ProtoFile) GetModelName() string {
	return strings.ToLower(strings.Replace(p.ModelName, " ", "_", -1))
}

type ProtoMessage struct {
	CommentList     []string
	Name            string
	ProtoFieldList  []*ProtoField
	maxProtoTypeLen int
	maxProtoNameLen int
}

type ProtoRPC struct {
	ProtoMessage
	RPCReturned string // the sent Message is returned
}

func (p *ProtoRPC) GetRPCReturnd() string {
	var packageName = "botproto."
	switch p.RPCReturned {
	case "array of the sent Messages":
		return "[]*" + packageName + "Message"
	case "":
		return ""
	case "True":
		return "*bool"
	case "String":
		return "*string"
	case "Int":
		return "*int32"
	case "Array of ChatMember":
		return "[]*" + packageName + "ChatMember"
	case "Array of GameHighScore":
		return "[]*" + packageName + "GameHighScore"
	default:
		return "*" + packageName + p.RPCReturned
	}
}

func (p *ProtoRPC) GetNewRPCReturndStr() string {
	var packageName = "botproto."
	switch p.RPCReturned {
	case "array of the sent Messages", "Array of ChatMember", "Array of GameHighScore":
		return "var rst = make(" + p.GetRPCReturnd() + ", 0)"
	case "":
		return ""
	case "True":
		return "var rst = new(bool)"
	case "String":
		return "var rst = new(string)"
	case "Int":
		return "var rst = new(int32)"
	default:
		return "var rst = new(" + packageName + p.RPCReturned + ")"
	}
}

type ProtoField struct {
	ProtoType string
	ProtoName string
	Required  string
	Comment   string
}

const (
	INDEX_ProtoName     = 1
	INDEX_ProtoType     = 2
	INDEX_ProtoRequired = 3
	INDEX_protoComment  = 4
)

func (p *ProtoField) GetProtoType() string {
	return getProtoType(p.ProtoType)
}

func getProtoType(protoType string) string {
	switch protoType {
	case "Integer":
		return "int32"
	case "String":
		return "string"
	case "Boolean", "True":
		return "bool"
	case "Float", "Float number":
		return "int64"
	case "InputFile or String":
		return "string"
	case "Integer or String":
		return "int32"
	case "InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply":
		return "ReplyMarkup"
	case "Array of InputMediaPhoto and InputMediaVideo":
		return "repeated InputMedia"
	default:
		if strings.HasPrefix(protoType, "Array of ") {
			return "repeated " + getProtoType(strings.Replace(protoType, "Array of ", "", 2))
		}
		return strings.Replace(protoType, "Array of ", "", 1)
	}
}

func (p *ProtoField) GetProtoTypeLen() int {
	return len(p.GetProtoType())
}

func (p *ProtoField) GetProtoName() string {
	return p.ProtoName
}

func (p *ProtoField) GetGOName() string {
	return generate.CamelCase(p.ProtoName)
}

func (p *ProtoField) GetProtoNameLen() int {
	return len(p.GetProtoName())
}

func (p *ProtoFile) AddMessageComment(isMethod bool, comment ...string) {
	if (isMethod && p.curRPC == nil) || (!isMethod && p.curMessage == nil) {
		log.Println("// ", comment)
		return
	}
	if isMethod {
		for _, one := range comment {
			for _, exp := range expList {
				strs := exp.FindAllStringSubmatch(one, 1)
				if len(strs) > 0 && len(strs[0]) > 0 {
					p.curRPC.RPCReturned = strs[0][1]
					break
				}
			}
		}
		p.curRPC.CommentList = append(p.curRPC.CommentList, comment...)
	} else {
		p.curMessage.CommentList = append(p.curMessage.CommentList, comment...)
	}
}

func (p *ProtoFile) IsMethod(name string) bool {
	if name[0] > 'a' && name[0] < 'z' {
		return true
	}
	return false
}

func (p *ProtoFile) AddMessage(isMethod bool, name string) {
	if isMethod {
		p.ProtoRPCList = append(p.ProtoRPCList, &ProtoRPC{})
		p.curRPC = p.ProtoRPCList[len(p.ProtoRPCList)-1]
		p.curRPC.Name = name
	} else {
		p.ProtoMessageList = append(p.ProtoMessageList, &ProtoMessage{})
		p.curMessage = p.ProtoMessageList[len(p.ProtoMessageList)-1]
		p.curMessage.Name = name
	}
}

func (p *ProtoFile) AddField(isMethod bool) {
	if isMethod {
		p.curRPC.ProtoFieldList = append(p.curRPC.ProtoFieldList, &ProtoField{})
	} else {
		p.curMessage.ProtoFieldList = append(p.curMessage.ProtoFieldList, &ProtoField{})
	}
}

func (p *ProtoFile) AddFieldByIndex(isMethod bool, index int, str string) {
	var curField *ProtoField
	if isMethod {
		curField = p.curRPC.ProtoFieldList[len(p.curRPC.ProtoFieldList)-1]
	} else {
		curField = p.curMessage.ProtoFieldList[len(p.curMessage.ProtoFieldList)-1]
		if index == INDEX_ProtoRequired {
			index++
		}
	}

	switch index {
	case INDEX_ProtoName:
		curField.ProtoName = str
	case INDEX_ProtoType:
		curField.ProtoType = str
	case INDEX_ProtoRequired:
		curField.Required = str
	case INDEX_protoComment:
		curField.Comment = str
	default:
		log.Println("AddFieldByIndex unknown index", index)
	}

	if isMethod {
		switch index {
		case INDEX_ProtoName:
			if p.curRPC.maxProtoNameLen < len(str) {
				p.curRPC.maxProtoNameLen = len(str)
			}
		case INDEX_ProtoType:
			if l := len(curField.GetProtoType()); p.curRPC.maxProtoTypeLen < l {
				p.curRPC.maxProtoTypeLen = l
			}
		default:
		}
	} else {
		switch index {
		case INDEX_ProtoName:
			if p.curMessage.maxProtoNameLen < len(str) {
				p.curMessage.maxProtoNameLen = len(str)
			}
		case INDEX_ProtoType:
			if l := len(curField.GetProtoType()); p.curMessage.maxProtoTypeLen < l {
				p.curMessage.maxProtoTypeLen = l
			}
		default:
		}
	}
}

func (p *ProtoFile) SetMessageName(name string) {
	p.curMessage.Name = name
}
