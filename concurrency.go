package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println("Num CPUs", runtime.NumCPU())
	fmt.Println("Num goroutines", runtime.NumGoroutine())
	wg.Add(3)
	go foo()
	go boo()
	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Println("Anon func", i)
		}
		wg.Done()
	}()

	fmt.Println("Num CPUs", runtime.NumCPU())
	fmt.Println("Num goroutines", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()
}

func boo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("boo", i)
	}
}
