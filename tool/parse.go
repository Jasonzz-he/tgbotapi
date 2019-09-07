package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/Jasonzz-he/golanglib/generate"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"golang.org/x/net/html/atom"
)

/*
   创建日期：2019/9/4
   修改日期：2019/9/4
   作者：Jason
*/

func H3HTMLCallback(e *colly.HTMLElement) {
	var (
		p        = NewProtoFile(e.Text)
		isMehtod bool
	)
	switch e.Text {
	case "Available types", "Games", "Stickers", "Payments", "Telegram Passport", "Available methods", "Getting updates", "Updating messages", "Inline mode":
		se := e.DOM.Next()
	FORLOOP:
		for len(se.Nodes) > 0 {
			switch se.Nodes[0].DataAtom {
			case atom.H3:
				break FORLOOP
			case atom.P:
				p.AddMessageComment(isMehtod, se.Text())
			case atom.Blockquote:
				p.AddMessageComment(isMehtod, strings.TrimLeft(se.First().Text(), "\n"))
			case atom.H4:
				switch se.Text() {
				case "Inline mode objects", "Sending files", "Inline mode methods", "Formatting options", "InputMedia":
					// don`t parse
				default:
					isMehtod = p.IsMethod(se.Text())
					p.AddMessage(isMehtod, se.Text())
				}
			case atom.Table, atom.Ul: // Ul: InputMedia 记录文字
				se.Each(func(i int, selection *goquery.Selection) {
					if selection.Nodes[0].FirstChild != nil {
						for c := selection.Nodes[0].FirstChild; c != nil; c = c.NextSibling {
							switch c.DataAtom {
							case atom.Thead:
							case atom.Tbody:
								for c := c.FirstChild; c != nil; c = c.NextSibling {
									switch c.DataAtom {
									case atom.Tr:
										p.AddField(isMehtod)
										var index int
										for c := c.FirstChild; c != nil; c = c.NextSibling {
											switch c.DataAtom {
											case atom.Td:
												index++
												var val string
												for c := c.FirstChild; c != nil; c = c.NextSibling {
													switch c.DataAtom {
													case 0:
														val += c.Data
													case atom.A, atom.Code, atom.Em, atom.Strong:
														val += c.FirstChild.Data
													case atom.Br: // 换行不需要

													default:
														if isMehtod {
															log.Println("00000000", index, isMehtod, p.curRPC.Name, c.Data, c.DataAtom.String())
														} else {
															log.Println("00000000", index, isMehtod, p.curMessage.Name, len(p.curMessage.ProtoFieldList), p.curMessage.ProtoFieldList[len(p.curMessage.ProtoFieldList)-1], c.Data, c.DataAtom.String())
														}
													}
												}
												p.AddFieldByIndex(isMehtod, index, val)
											case 0:

											default:
												log.Printf("=========%s %s %#v", c.DataAtom.String(), c.Data, c.FirstChild)
											}
										}
									case atom.Atom(0):
									default:
										log.Println("-----", c.DataAtom.String())
									}
								}
							case atom.Li:
								switch c.FirstChild.DataAtom {
								case atom.A, atom.Strong:
									p.AddMessageComment(isMehtod, c.FirstChild.FirstChild.Data)
								case 0:

								default:
									log.Println("-------=========", c.FirstChild.DataAtom.String(), isMehtod)
								}

							case atom.Atom(0):
							default:
								log.Printf("----------- %T", c.DataAtom)
							}
						}
					}
				})
			case atom.H6, atom.Pre:
				// H6: Markdown style 忽略
			case atom.Div:
				// LoginUrl 图片 忽略

			default:
				if isMehtod {
					log.Printf("-----%s %v %d %#v", se.Nodes[0].DataAtom.String(), p.curRPC.Name, len(se.Nodes), se.Nodes[0])
				} else {
					log.Println("-----", se.Nodes[0].DataAtom.String())
				}
			}
			se = se.Next()
		}

	default:
		log.Println(e.Text)
	}
	//log.Println(e.DOM.Next().Text())
}

func getTGBotApiGenerate(dirName, fileName string) *generate.Generate {
	g := generate.New(dirName, fileName)
	g.P("syntax = ", g.Symbol('"', "proto3"), ";").P()
	g.P("package botproto;").P()
	return g
}

func getTGBotApiMethodGoFileGenerate(dirName, fileName string) *generate.Generate {
	g := generate.New(dirName, fileName)
	g.P("package tgbotapi").P()
	g.P("import ", g.Symbol('"', "github.com/Jasonzz-he/tgbotapi/botproto")).P()
	return g
}

func getTGBotApiMethodGoTestFileGenerate(dirName, fileName string) *generate.Generate {
	g := generate.New(dirName, fileName)
	g.P("package tgbotapi").P()
	g.P("import (").InFunc(")", func() {
		g.P(g.Symbol('"', "log"))
		g.P(g.Symbol('"', "testing"))
		g.P(g.Symbol('"', "github.com/Jasonzz-he/tgbotapi/botproto"))
		g.P(g.Symbol('"', "github.com/smartystreets/assertions"))
	}).P()
	return g
}

var (
	expList = []*regexp.Regexp{
		regexp.MustCompile(`On success, the sent ([a-zA-Z_]+) is returned.`),
		regexp.MustCompile(`On success, returns an (Array of [a-zA-Z_]+) objects`),
		regexp.MustCompile(`On success, an (array of the sent [a-zA-Z_]+) is returned.`),
		regexp.MustCompile(`Returns basic information about the bot in form of a ([a-zA-z_]+) object.`),
		regexp.MustCompile(`Returns a ([a-zA-z_]+) object on success.`),
		regexp.MustCompile(` ([a-zA-z_]+) on success.`),
	}
)

func Write(ps []*ProtoFile) {
	var outFilePath = "../botproto"
	genProtoMessage(outFilePath, ps)
	genProtoRPC(outFilePath, ps)

	outFilePath = "../"
	genMethodFile(outFilePath, ps)
}

func genMethodFile(outFilePath string, ps []*ProtoFile) {
	for _, p := range ps {
		if len(p.ProtoRPCList) > 0 {
			g := getTGBotApiMethodGoFileGenerate("", "gen_"+p.GetModelName()+".go")
			testG := getTGBotApiMethodGoTestFileGenerate("", "gen_"+p.GetModelName()+"_test.go")
			for _, protoRPC := range p.ProtoRPCList {
				var (
					inParamStr  = "in *botproto." + strings.Title(protoRPC.Name)
					inParamName = "in"
				)
				if len(protoRPC.ProtoFieldList) == 0 {
					inParamStr = ""
					inParamName = "nil"
				}
				if protoRPC.RPCReturned == "" {
					g.P("func (b *Bot) ", strings.Title(protoRPC.Name), "(", inParamStr, ") error {").InFunc("}", func() {
						g.P("err := b.Post(", g.Symbol('"', protoRPC.Name), ", ", inParamName, ", nil)")
						g.P("return err")
					}).P()

					g.P("func  ", strings.Title(protoRPC.Name), "(", inParamStr, ") error {").InFunc("}", func() {
						if len(protoRPC.ProtoFieldList) == 0 {
							g.P("return globalBot.", strings.Title(protoRPC.Name), "()")
						} else {
							g.P("return globalBot.", strings.Title(protoRPC.Name), "(in)")
						}
					}).P()
				} else {
					g.P("func (b *Bot) ", strings.Title(protoRPC.Name), "(", inParamStr, ") (", protoRPC.GetRPCReturnd(), ", error) {").InFunc("}", func() {
						g.P(protoRPC.GetNewRPCReturndStr())
						g.P("err := b.Post(", g.Symbol('"', protoRPC.Name), ", ", inParamName, ", rst)")
						g.P("return rst, err")
					}).P()

					g.P("func  ", strings.Title(protoRPC.Name), "(", inParamStr, ") (", protoRPC.GetRPCReturnd(), ", error) {").InFunc("}", func() {
						if len(protoRPC.ProtoFieldList) == 0 {
							g.P("return globalBot.", strings.Title(protoRPC.Name), "()")
						} else {
							g.P("return globalBot.", strings.Title(protoRPC.Name), "(in)")
						}
					}).P()
				}

				// test go file
				testG.P("func TestBot_", strings.Title(protoRPC.Name), "(t *testing.T) {").InFunc("}", func() {
					if protoRPC.RPCReturned == "" {
						if len(protoRPC.ProtoFieldList) > 0 {
							testG.P("err := ", strings.Title(protoRPC.Name), "(&botproto.", strings.Title(protoRPC.Name), "{").InFunc("})", func() {
								genTestGoParam(testG, protoRPC.ProtoFieldList)
							})
						} else {
							testG.P("err := ", strings.Title(protoRPC.Name), "()")
						}
					} else {
						if len(protoRPC.ProtoFieldList) > 0 {
							testG.P("rst, err := ", strings.Title(protoRPC.Name), "(&botproto.", strings.Title(protoRPC.Name), "{").InFunc("})", func() {
								genTestGoParam(testG, protoRPC.ProtoFieldList)
							})
						} else {
							testG.P("rst, err := ", strings.Title(protoRPC.Name), "()")
						}
					}
					testG.P("log.Println(assertions.ShouldBeNil(err))")
					if protoRPC.RPCReturned != "" {
						testG.P("if nil == err {").InFunc("}", func() {
							testG.P("log.Printf(\"%#v\", rst)")
						})
					}
				}).P()
			}
			g.WriteFile(outFilePath)
			testG.WriteFile(outFilePath)
		}
	}
}

func genTestGoParam(g *generate.Generate, fieldList []*ProtoField) {
	for _, one := range fieldList {
		var val string
		switch one.GetProtoType() {
		case "string":
			val = "\"\""
		case "int32", "int64":
			val = "0"
		case "bool":
			val = "false"
		default:
			val = "nil"
		}
		g.P(one.GetGOName(), " : ", val, ", // ", one.GetProtoType())
	}
}

func genProtoMessage(outFilePath string, ps []*ProtoFile) {
	g := getTGBotApiGenerate("", "tgbotapi.proto")
	for _, p := range ps {
		for _, comment := range p.Comment {
			g.P("// ", comment)
		}
		for _, protoMessage := range p.ProtoMessageList {
			genFieldList(g, protoMessage)
		}
	}
	var extra = `
message ReplyMarkup {
	oneof ReplyMarkup {
		InlineKeyboardMarkup InlineKeyboardMarkup = 1;
		ReplyKeyboardMarkup  ReplyKeyboardMarkup  = 2;
		ReplyKeyboardRemove  ReplyKeyboardRemove  = 3;
		ForceReply			 ForceReply			  = 4;
	}
}

// This object represents the content of a media message to be sent. It should be one of
// InputMediaAnimation
// InputMediaDocument
// InputMediaAudio
// InputMediaPhoto
// InputMediaVideo
message InputMedia {
	oneof InputMedia {
		InputMediaAnimation InputMediaAnimation = 1;
		InputMediaDocument  InputMediaDocument  = 2;
		InputMediaAudio     InputMediaAudio     = 3;
		InputMediaPhoto     InputMediaPhoto     = 4;
		InputMediaVideo     InputMediaVideo     = 5;
	}
}
`
	g.P(extra)
	g.WriteFile(outFilePath)
}

func genProtoRPC(outFilePath string, ps []*ProtoFile) {
	for _, p := range ps {
		if len(p.ProtoRPCList) > 0 {
			g := getTGBotApiGenerate("", strings.Join([]string{p.GetModelName(), "method", "proto"}, "."))
			g.P("import \"tgbotapi.proto\";").P()
			for _, protoRPC := range p.ProtoRPCList {
				g.P("// return ", protoRPC.RPCReturned)
				genFieldList(g, &protoRPC.ProtoMessage)
			}
			g.WriteFile(outFilePath)
		}
	}
}

func genFieldList(g *generate.Generate, protoMessage *ProtoMessage) {
	for _, protoMessageComment := range protoMessage.CommentList {
		g.P("// ", strings.Replace(protoMessageComment, "\n", "\n// ", -1))
	}
	g.P("message ", protoMessage.Name, " {").InFunc("}", func() {
		for index, field := range protoMessage.ProtoFieldList {
			g.P(field.GetProtoType(),
				strings.Repeat(" ", 1+protoMessage.maxProtoTypeLen-field.GetProtoTypeLen()),
				field.GetProtoName(),
				strings.Repeat(" ", protoMessage.maxProtoNameLen-field.GetProtoNameLen()),
				" = ", index+1, "; // ", field.Required, " ", field.Comment)
		}
	}).P()
}
