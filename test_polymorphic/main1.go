package main

import "fmt"

// declare a function type
type Operation func(int, int) int

func Add(a int, b int) int {
	return a + b
}

func Multiply(a int, b int) int {
	return a * b
}

func caculate(a int, b int, op Operation) int {
	return op(a, b)
}

func main() {
	fmt.Println(caculate(2, 3, Add))
	fmt.Println(caculate(2, 3, Multiply))
}
