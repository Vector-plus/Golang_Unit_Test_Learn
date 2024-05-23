package example

import "fmt"

func Mult(a, b int) int {
	x := Mul_Func(a, b)
	fmt.Printf("the result of (a+b) = %d\n", x)
	return x
}

var Mul_Func = add

func add(a, b int) int {
	return a + b
}
