package main

import "fmt"

func main() {
	si := []int{1, 2, 3, 4}
	for i, v := range takeInt(si...) {
		fmt.Println(i, v)
	}
}

func takeInt(x ...int) []int {
	return x
}
func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
