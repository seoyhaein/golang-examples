package main

type nums []int

var sum int

func inputnums() int {
	ns := nums{1, 2, 3}
	for _, n := range ns {
		sum = sum + n
	}
	return sum
}
