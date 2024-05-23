package httptest

import (
	"encoding/json"
	"fmt"
	"gotest/dv1/example"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type Resp struct {
	Code string
	Data interface{}
	Msg  string
}

var msgData = &Resp{
	Code: "200",
	Data: &example.User{UserName: "fht", Password: "123456", Age: 12},
	Msg:  "ok",
}

func HttpSrv(w http.ResponseWriter, r *http.Request) {
	//在该方法中还能处理请求路径
	// r.URL.EscapedPath() != "/userInfo"
	// 获取请求参数：r.ParseForm，r.Form.Get("addr")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		data, _ := json.Marshal(msgData)
		w.Write(data)
	}
}

func TestHttpSrv(t *testing.T) {
	Convey("测试httpSrv", t, func() {
		// 通过httptest.NewServer创建了一个测试的http server,写自己的HandlerFunc函数，处理请求。
		ts := httptest.NewServer(http.HandlerFunc(HttpSrv))
		defer ts.Close()

		resp, err := example.GetUserInfo(ts.URL)
		fmt.Println(resp)
		So(err, ShouldBeNil)
		param, _ := json.Marshal(msgData)
		So(resp, ShouldEqual, string(param))
	})
}
