package main

import "fmt"

func main() {
	fmt.Println(bar()())
	fmt.Println(func(a int, b int) int {
		return a + b
	}(2, 3))

}

func bar() func() int {
	return func() int {
		return 451
	}
}
