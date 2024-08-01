package fanxing

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTT(t *testing.T) {
	Convey("测试泛型函数", t, func() {
		n := 10000
		arry := make([]string, n)
		arry1 := make([]any, n)
		for i := 0; i < n; i++ {
			arry1[i] = "kkk"
		}
		start_tim := time.Now()
		for i := 0; i < n; i++ {
			IsExist(arry1, "qw")
		}
		fmt.Printf("泛型花费时间：%v\n", time.Since(start_tim))

		start_tim = time.Now()
		for i := 0; i < n; i++ {
			IsEt(arry, "as")
		}
		fmt.Printf("非泛型花费时间：%v\n", time.Since(start_tim))

	})
}

// func IsExist[T string | int | int32 | byte](arry []T, item T) bool {
// 	for _, v := range arry {
// 		if v == item {
// 			return true
// 		}
// 	}
// 	return false
// }

func IsExist(arry []any, item any) bool {
	for _, v := range arry {
		if v == item {
			return true
		}
	}
	return false
}

func IsEt(arry []string, item string) bool {
	for _, v := range arry {
		if v == item {
			return true
		}
	}
	return false
}
