package main

import (
	"bytes"
	"fmt"
)

type (
	Opt  func(s string) string
	Opts []Opt
	Nums []int
)

func (opts Opts) printOpts() string {
	var b bytes.Buffer
	for _, opt := range opts {
		s := opt("Printed")
		b.WriteString(s)
	}
	return b.String()
}

func Sum(nums ...int) int {
	ns := Nums(nums)
	var sum int
	for _, n := range ns {
		sum = sum + n
	}
	return sum
}

func New(opts ...Opt) string {
	s := Opts(opts).printOpts()
	return s
}

func main() {
	fmt.Println("slice tester")
	n := Sum(1, 2, 3, 4)

	fmt.Println(n)

	fn := func(s string) string {
		var b bytes.Buffer
		b.WriteString("hello")
		b.WriteString(s)
		return b.String()
	}
	fmt.Println(New(fn))

	fmt.Println("parameter test")
	number := Caller(100)

	fmt.Println("n = ", number)

	fmt.Println("slice test")
	n1 := inputnums()
	fmt.Println(n1)
}
