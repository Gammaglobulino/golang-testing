package main

import "fmt"

func main() {
	c := make(chan int)
	go sendc(c)
	fmt.Println(receivec(c))
}

func sendc(c chan<- int) {
	c <- 42
}
func receivec(c <-chan int) int {
	return <-c

}
