package stubtest

import (
	"gotest/dv1/example"
	"testing"

	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStub(t *testing.T) {
	Convey("测试stub替换函数或全局变量", t, func() {
		//使用stub需要在源代码中定义一个全局变量，每次替换时会自动更改其值，使用完之后，Reset()方法自动还原该值
		stub := gostub.Stub(&example.Mul_Func, func(a, b int) int {
			return a - b
		})
		defer stub.Reset()

		x := example.Mult(1, 2)
		So(x, ShouldEqual, -1)
	})
}
