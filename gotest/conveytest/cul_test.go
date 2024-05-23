package conveytest

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// 命令行查看测试结果（查看覆盖率） go test -v -cover
// web-UI查看测试结果（设置端口） goconvey -port 8099
func TestCulculate(t *testing.T) {
	//新增测试
	Convey("测试Add函数", t, func() {
		//断言判断结果是否正确
		So(Add(1, 2), ShouldEqual, 3)
	})

	Convey("测试Subt函数", t, func() {
		//断言函数判断
		So(Subt(1, 3), ShouldEqual, -2)
		So(Subt(10, 5), ShouldEqual, 5)
	})

	Convey("测试Mult函数", t, func() {
		//断言判断
		So(Mult(1, 2), ShouldEqual, 2)
		So(Mult(4, 6), ShouldEqual, 24)
	})

	Convey("测试Div函数", t, func() {
		//嵌套使用convey，只有最外层convey需要传入*testing.T参数
		Convey("除数为0", func() {
			_, err := Div(1, 0)
			So(err, ShouldNotBeNil)
			_, err = Div(7, 0)
			So(err, ShouldNotBeNil)
			_, err = Div(1999999, 0)
			So(err, ShouldNotBeNil)
		})
		Convey("除数不为0", func() {
			re, err := Div(4, 2)
			So(err, ShouldBeNil)
			So(re, ShouldEqual, 2)

			re, err = Div(5, 2)
			So(err, ShouldBeNil)
			So(re, ShouldEqual, 2.5)

			re, err = Div(20, 2)
			So(err, ShouldBeNil)
			So(re, ShouldEqual, 10)
		})
	})
}
