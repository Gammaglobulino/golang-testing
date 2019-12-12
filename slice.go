package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6}

	for i, v := range sl {
		fmt.Printf("%d - value: %v - capacity:%d - lenght:%d \n", i, v, cap(sl), len(sl))
	}
	for i := 0; i < 10; i++ {
		sl = append(sl, i)
		fmt.Printf("value: %v - capacity:%d - lenght:%d \n", i, cap(sl), len(sl))
	}
}
