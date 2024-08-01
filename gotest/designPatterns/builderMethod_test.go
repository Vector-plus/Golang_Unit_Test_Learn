package designpatterns

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBuilder(t *testing.T) {
	Convey("测试生成器", t, func() {
		intel, _ := getBuilder("Intel", "Kinston")
		amd, _ := getBuilder("AMD", "SAMSUNG")
		direct := newDirector(intel)
		intelComp := direct.makeComputer()

		fmt.Println(intelComp)
		So(intelComp.CPUType, ShouldEqual, "Intel")

		direct.setBuilder(amd)
		amdComp := direct.makeComputer()

		fmt.Println(amdComp)
		So(amdComp.CPUType, ShouldEqual, "AMD")
	})
}
