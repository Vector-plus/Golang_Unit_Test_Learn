package designpatterns

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIterator(t *testing.T) {
	Convey("测试迭代器", t, func() {
		userCol := &ListCollection{
			userList: make([]any, 0),
		}

		for i := int64(0); i < 10; i++ {
			userCol.userList = append(userCol.userList, &User{name: "fht--" + strconv.FormatInt(i, 10), age: i + 10})
		}

		iterator := userCol.getIterator()

		for iterator.hasNext() {
			user, err := iterator.getNext()
			So(err, ShouldBeNil)
			fmt.Println(user)
		}

	})

}
