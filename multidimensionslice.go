package main

import "fmt"

func main() {
	sl := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	for _, v := range sl {
		for _, x := range v {
			fmt.Print(x)
		}
		fmt.Println()
	}
}
