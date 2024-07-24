package designpatterns

import (
	"fmt"
	"strconv"
)

//适配器，用于统一不同的接口，相当于一个中转处理，让接口调用之间兼容

type AppSend interface {
	SendMsgNum(num string)
}

type Client struct {
}

func (c *Client) Send(app AppSend, msg string) {
	app.SendMsgNum(msg)
}

type Wechat struct {
}

func (app *Wechat) SendMsgssssssss(nums int64) {
	fmt.Println("I am Wechat ----> ", nums)
}

func (app *Wechat) SendMsgNum(num string) {
	//适配器，进行转换
	msg, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		panic("输入数字错误")
	}
	app.SendMsgssssssss(msg)
}

type QQ struct {
}

func (app *QQ) SendMsggggggggs(nums int32) {
	fmt.Println("I am QQ -----> ", nums)
}

func (app *QQ) SendMsgNum(num string) {
	//适配器，进行转换，调用方法
	msg, err := strconv.ParseInt(num, 10, 32)
	if err != nil {
		panic("输入参数异常")
	}
	app.SendMsggggggggs(int32(msg))
}
