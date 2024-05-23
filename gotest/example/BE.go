package example

import (
	"errors"
	"fmt"
	"time"
)

// 该struct没有interface的实现，故不能使用monkey的方式。
type UnitT2 struct {
}

func (u UnitT2) WiteAndSend2(txt string, n int) (string, error) {
	ans, err := u.FA2(txt)

	if err != nil {
		return "", err
	}

	x, err := u.FB2(n)
	if err != nil {
		return "", err
	}

	y, err := u.FB2(x)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s === %d -----> %d\n", ans, x, y), nil
}

func (u *UnitT2) FA2(msg string) (string, error) {
	if msg == "" {
		return "", errors.New("参数为空")
	}
	return fmt.Sprintf("%s.....%s", msg, time.Now().Format(time.RFC1123)), nil
}

func (u *UnitT2) FB2(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("测试参数错误")
	}
	if n <= 2 {
		return 1, nil
	}
	x1, _ := u.FB2(n - 1)
	x2, _ := u.FB2(n - 2)
	return x1 + x2, nil
}
