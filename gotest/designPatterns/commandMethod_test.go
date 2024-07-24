package designpatterns

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCommand(t *testing.T) {
	Convey("测试命令", t, func() {
		tv := &TV{}

		tvOn := &OnCommand{Device: tv}
		tvOff := &OffCommand{Device: tv}

		tvOn.Excute()
		tvOff.Excute()

	})
}
