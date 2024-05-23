package example

import (
	"errors"
	"fmt"
	"time"
)

type IA interface {
	//给txt添加时间戳
	FA(msg string) (string, error)
	//获取第n项斐波那契数
	FB(n int) (int, error)
}

type UnitT struct {
	ia IA
}

func NewIA(ia IA) *UnitT {
	return &UnitT{ia: ia}
}

func (u UnitT) WiteAndSend(txt string, n int) (string, error) {
	ans, err := u.ia.FA(txt)

	if err != nil {
		return "", err
	}

	x, err := u.ia.FB(n)
	if err != nil {
		return "", err
	}

	y, err := u.ia.FB(x)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s === %d -----> %d\n", ans, x, y), nil
}

func (u *UnitT) FA(msg string) (string, error) {
	if msg == "" {
		return "", errors.New("参数为空")
	}
	return fmt.Sprintf("%s.....%s", msg, time.Now().Format(time.RFC1123)), nil
}

func (u *UnitT) FB(n int) (int, error) {
	if n < 0 {
		return -1, errors.New("测试参数错误")
	}
	if n <= 2 {
		return 1, nil
	}
	x1, _ := u.FB(n - 1)
	x2, _ := u.FB(n - 2)
	return x1 + x2, nil
}
