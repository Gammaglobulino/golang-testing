package main

import "fmt"

func main() {
	intch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 9; i++ {
			intch <- i
		}
		done <- true
	}()

	go func() {
		<-done
		close(intch)
	}()

	for v := range intch {
		fmt.Println(v)
	}
}
