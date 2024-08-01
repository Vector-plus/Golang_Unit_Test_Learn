package designpatterns

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStrategy(t *testing.T) {
	Convey("测试策略模式", t, func() {

		//新建策略
		ahead1 := &addOneStep{}
		ahead2 := &addTwoStep{}
		ahead3 := &addThreeStep{}

		//新建执行对象
		run := newRunStep(0, ahead1)

		So(run.n, ShouldEqual, 0)

		//执行策略
		run.Run()

		So(run.n, ShouldEqual, 1)

		run.setStrategy(ahead2)
		run.Run()

		So(run.n, ShouldEqual, 3)

		run.setStrategy(ahead3)
		run.Run()

		So(run.n, ShouldEqual, 6)
	})
}
