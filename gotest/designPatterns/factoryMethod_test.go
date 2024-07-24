package designpatterns

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMakComputer(t *testing.T) {
	Convey("测试生产工厂模式", t, func() {
		invida, err := getComputer("INVIDA")
		gpu := invida.makeGpu()
		cache := invida.makeCache()
		cpu := invida.makeCpu()
		gpu.setName("invida")
		cache.setSize(32 * 1024 * 1024)
		cpu.setName("intel")
		So(err, ShouldBeNil)

	})
}
