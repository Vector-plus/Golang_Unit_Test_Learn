package designpatterns

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdapter(t *testing.T) {
	Convey("测试适配器", t, func() {
		client := &Client{}
		wechat := &Wechat{}
		qq := &QQ{}

		client.Send(qq, "64")
		client.Send(wechat, "32")
	})
}
