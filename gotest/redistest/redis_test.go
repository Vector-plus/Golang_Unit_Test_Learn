package redistest

import (
	"gotest/dv1/example"
	"testing"

	"github.com/redis/go-redis/v9"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/alicebob/miniredis/v2"
)

// miniredis是一个纯go实现的用于单元测试的redis server。它是一个简单易用的、基于内存的redis替代品，
// 它具有真正的TCP接口，你可以把它当成是redis版本的net/http/httptest。
// 当我们为一些包含Redis操作的代码编写单元测试时就可以使用它来mock Redis操作。
func TestMiniRedis(t *testing.T) {

	Convey("测试miniRedis", t, func() {
		//使用minireids新建一个redis
		rs, err := miniredis.Run()
		So(err, ShouldBeNil)
		defer rs.Close()

		key := "2024-05-20"

		//预置数据到redis中
		_, err = rs.SAdd(key, "test---miniredis")
		So(err, ShouldBeNil)

		reclient := redis.NewClient(&redis.Options{Addr: rs.Addr()})
		rdb := &example.Rdb{Rdb: reclient}

		ans, err := rdb.GetAllUser(key)

		So(err, ShouldBeNil)
		So(ans[0], ShouldEqual, "test---miniredis")

	})

}
