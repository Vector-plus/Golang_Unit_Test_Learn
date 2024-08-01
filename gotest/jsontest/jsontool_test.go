package jsontest

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type Student struct {
	Name string `json:"name"`
	Flag bool   `json:"flag"`
}

func TestMashll(t *testing.T) {
	Convey("测试json解析", t, func() {
		// str := "[{"name":"fht","flag":true},{"name":"fh","flag":false}]"
		stus := make([]*Student, 0)

		stus = append(stus, &Student{Name: "fht", Flag: true})
		stus = append(stus, &Student{Name: "fh", Flag: false})

		strStu, err := json.Marshal(stus)
		So(err, ShouldBeNil)
		fmt.Println(string(strStu))

		var ss []*Student
		err = json.Unmarshal(strStu, &ss)
		So(err, ShouldBeNil)
		fmt.Println(*ss[0])
		fmt.Println(*ss[1])
	})

}
