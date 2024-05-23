package monkeytest

import (
	"gotest/dv1/example"
	"reflect"
	"testing"

	"bou.ke/monkey"
	. "github.com/smartystreets/goconvey/convey"
)

// 使用monkey方式，直接替换对象中的方法，其原理为采用热补丁的方式替换二进制文件中的代码，在执行时直接跳转到桩实现
// 使用时需要使用 -gcflags=-l 关闭内联优化（跳转到正确的执行文件，该方式影响了源文件）
// 该方式不能用于多线程的测试
// 运行指令 go test -run=TestUser -v -gcflags=-l
func TestWiteAndSend2(t *testing.T) {
	Convey("monkey测试WithAndSend2方法", t, func() {

		unitT2 := &example.UnitT2{}

		//为对象方法进行打桩
		monkey.PatchInstanceMethod(reflect.TypeOf(unitT2), "FA2", func(*example.UnitT2, string) (string, error) {
			return "test----2024-05-20", nil
		})

		monkey.PatchInstanceMethod(reflect.TypeOf(unitT2), "FB2", func(e *example.UnitT2, n int) (int, error) {
			return n + 1, nil
		})

		//直接调用测试方法，此时该对象的FA2和FB2已经被替换
		ans, err := unitT2.WiteAndSend2("test", 1)

		So(err, ShouldBeNil)
		So(ans, ShouldEqual, "test----2024-05-20 === 2 -----> 3\n")

	})
}
