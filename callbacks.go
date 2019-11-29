package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(command(sum, x))
}
func command(f func([]int) int, x []int) int {
	return f(x)
}

func sum(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
