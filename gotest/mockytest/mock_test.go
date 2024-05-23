package mockytest

import (
	"fmt"
	"gotest/dv1/example"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

// 测试go-Mock打桩，利用interface多态的特点生成测试对象，实现替换interface中的某些方法
// 生成对象指令（mockgen插件） mockgen -source=./mail.go -destination=../mocktest/mock_mockdemo.go -package=mocktest Mail
func TestMock(t *testing.T) {

	Convey("测试example.WiteAndSend(string,int)方法", t, func() {
		//获取mockctl，它代表mock生态系统中的顶级控件。定义了mock对象的范围、生命周期和期待值。另外它在多个goroutine中是安全的
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()

		//生成需要mock的实例
		mockIA := NewMockIA(mockCtl)

		//声明给定的调用应按顺序进行
		gomock.InOrder(

			//三个步骤：EXPECT()返回一个允许调用者设置期望和返回值的对象。FA("test")是设置入参并调用 mock 实例中的方法。Return是设置先前调用的方法出参。简单来说，就是设置入参并调用，最后设置返回值
			mockIA.EXPECT().FA("test").Return("test-----2024-05-20", nil),
			mockIA.EXPECT().FB(2).Return(1, nil),
			mockIA.EXPECT().FB(1).Return(1, nil),
		)

		//生成实例对象
		unitT := example.NewIA(mockIA)

		//调用测试方法
		ans, err := unitT.WiteAndSend("test", 2)
		So(err, ShouldBeNil)
		So(ans, ShouldEqual, "test-----2024-05-20 === 1 -----> 1\n")
		fmt.Println(ans)
	})

}
