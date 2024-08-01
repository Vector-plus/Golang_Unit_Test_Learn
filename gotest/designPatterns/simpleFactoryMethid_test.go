package designpatterns

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSimpleFactory(t *testing.T) {
	Convey("测试工厂方法", t, func() {
		dev1, err := getcamerDev("GB281", "陌生人闯入报警")

		So(err, ShouldBeNil)
		fmt.Println(dev1)

		dev2, err := getcamerDev("AI360Pro", "婴儿哭啼报警")

		So(err, ShouldBeNil)
		fmt.Println(dev2)
	})
}
