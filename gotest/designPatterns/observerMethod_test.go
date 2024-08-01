package designpatterns

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestObserver(t *testing.T) {
	Convey("测试观察者模式", t, func() {
		observer := &Observe{
			consumers: make(map[string]Consumer),
			name:      "taobao",
		}

		consumer1 := &Subscribe{
			id: "fuhaitao1",
		}

		consumer2 := &Subscribe{
			id: "fuhaitao2",
		}

		consumer3 := &Subscribe{
			id: "fuhaitao3",
		}

		//注册消费者
		observer.register(consumer1.id, consumer1)
		observer.register(consumer2.id, consumer2)
		observer.register(consumer3.id, consumer3)

		observer.ExistProducts()

	})
}
