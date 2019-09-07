### tgbotapi
telegram Bot API  golang interface

### 使用案例
```go
NewBot("", "654375916:AAF0alaqi5jVp0nG9755g1iZc1y6nawWFKU").ShowURL(true)
rst, err := GetMe()
log.Println(assertions.ShouldBeNil(err))
if nil == err {
    log.Printf("%#v", rst)
}
```

其他接口使用请参照各个test文件

### 交流群
```
https://t.me/jason_telegramBot
```