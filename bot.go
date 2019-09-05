package tgbotapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
   创建日期：2019/8/30
   修改日期：2019/8/30
   作者：Jason
*/

var globalBot *Bot

type Bot struct {
	url     string // 访问的potato机器人url 默认：https://api.telegram.org/bot
	token   string // 机器人唯一标志 exp: 654375916:AAF0alaqi5jVp0nG9755g1iZc1y6nawWFKU
	showUrl bool
}

func init() {
	NewBot("", "654375916:AAF0alaqi5jVp0nG9755g1iZc1y6nawWFKU").ShowURL(true)
}

func NewBot(botUrl, botToken string) *Bot {
	if strings.TrimSpace(botUrl) == "" {
		botUrl = DefaultBotUrl
	}
	globalBot = &Bot{
		url:   botUrl,
		token: botToken,
	}
	return globalBot
}

func (b *Bot) ShowURL(ok bool) *Bot {
	b.showUrl = ok
	return b
}

func ShowURL(ok bool) *Bot {
	return globalBot.ShowURL(ok)
}

func (b *Bot) getHttpURL(funcName string) string {
	return strings.Join([]string{b.url + b.token, funcName}, "/")
}

func (b *Bot) Post(methodName string, param interface{}, rst interface{}) error {
	url := b.getHttpURL(methodName)
	if b.showUrl {
		p, _ := json.MarshalIndent(param, "", "\t")
		fmt.Println(url, contentType)
		fmt.Println(string(p))
	}
	resp, err := http.Post(url, "application/json", jsonMarshal(param))
	if nil != err {
		return err
	}
	defer resp.Body.Close()
	return getBotResponse(resp.Body, rst)

}

func getBotResponse(r io.Reader, result interface{}) error {
	var resp = newBotResponse(result)
	respbs, err := ioutil.ReadAll(r)
	if nil != err {
		return err
	}

	err = json.Unmarshal(respbs, resp)
	if nil != err {
		return err
	} else if !resp.Ok {
		err = &botError{
			ErrCode:    resp.ErrorCode,
			ErrMessage: resp.Description,
		}
	}
	return err
}

const (
	DefaultBotUrl = "https://api.telegram.org/bot"

	contentType = "application/json"
)

type botResponse struct {
	Ok          bool        `json:"ok"`
	ErrorCode   int32       `json:"error_code"`
	Result      interface{} `json:"result"`
	Description string      `json:"description"`
}

func newBotResponse(result interface{}) *botResponse {
	return &botResponse{
		Result: result,
	}
}

func jsonMarshal(param interface{}) io.Reader {
	if param == nil {
		return nil
	}
	bs, _ := json.Marshal(param)
	return bytes.NewBuffer(bs)
}

type botError struct {
	ErrCode    int32
	ErrMessage string
}

func (e *botError) Error() string {
	return fmt.Sprintf("bot errCode: %d errMessage: %s", e.ErrCode, e.ErrMessage)
}

func NewBotError(errCode int32, errMessage string) error {
	return &botError{
		ErrCode:    errCode,
		ErrMessage: errMessage,
	}
}
