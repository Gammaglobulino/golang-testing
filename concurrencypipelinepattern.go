package main

import "fmt"

func main() {
	c := generator(2, 3, 4)
	out := square(c)
	for v := range out {
		fmt.Println(v)
	}

}

//generates an variadic number of ints using goroutines for every single number generated
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

//pull and square a series of ints from an inboud channel_ return an outbound number with the results
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
