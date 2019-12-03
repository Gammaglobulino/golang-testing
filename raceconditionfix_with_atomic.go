package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var x int64
	fmt.Println("Num CPU:", runtime.NumCPU())
	fmt.Println("Num goroutines", runtime.NumGoroutine())

	var wb sync.WaitGroup
	wb.Add(3)
	go func() {
		atomic.AddInt64(&x, 1)
		runtime.Gosched()
		fmt.Println("go one")
		wb.Done()
	}()
	go func() {
		atomic.AddInt64(&x, 1)
		runtime.Gosched()
		fmt.Println("go two")
		wb.Done()

	}()
	go func() {
		atomic.AddInt64(&x, 1)
		runtime.Gosched()
		fmt.Println("go three")
		wb.Done()
	}()
	wb.Wait()
	fmt.Println("main x", x)
	fmt.Println("Num CPU:", runtime.NumCPU())
	fmt.Println("Num goroutines", runtime.NumGoroutine())

}
