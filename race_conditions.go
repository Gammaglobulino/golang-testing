package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	x := 0
	fmt.Println("Num CPU:", runtime.NumCPU())
	fmt.Println("Num goroutines", runtime.NumGoroutine())

	var wb sync.WaitGroup
	wb.Add(3)
	go func() {

		v := x
		v++
		runtime.Gosched()
		fmt.Println("go one")
		x = v
		wb.Done()

	}()
	go func() {

		v := x
		v++
		runtime.Gosched()
		fmt.Println("go two")
		x = v
		wb.Done()

	}()
	go func() {

		v := x
		v++
		fmt.Println("go three")
		runtime.Gosched()
		x = v
		wb.Done()

	}()
	wb.Wait()
	fmt.Println("main x", x)
	fmt.Println("Num CPU:", runtime.NumCPU())
	fmt.Println("Num goroutines", runtime.NumGoroutine())

}
