package main

import "fmt"

func main() {
	f := incr()
	fmt.Println(f())
	fmt.Println(f())
}

func incr() func() int {
	x := 0
	return func() int {
		x++
		return (x)
	}
}
