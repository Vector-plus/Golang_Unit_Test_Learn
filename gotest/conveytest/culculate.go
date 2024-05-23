package conveytest

import "errors"

// 两数相加函数
func Add(a, b int) int {
	return a + b
}

// 两数差值函数
func Subt(a, b int) int {
	return a - b
}

// 两数相乘函数
func Mult(a, b int) int {
	return a * b
}

// 两数相除函数
func Div(a, b float64) (re float64, err error) {
	if b == 0 {
		re = -1
		err = errors.New("除数为0")
	} else {
		re = a / b
	}
	return
}
