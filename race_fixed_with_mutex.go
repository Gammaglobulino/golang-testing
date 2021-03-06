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
	var mt sync.Mutex
	var wb sync.WaitGroup
	wb.Add(3)
	go func() {
		mt.Lock()
		v := x
		v++
		x = v
		mt.Unlock()
		runtime.Gosched()
		fmt.Println("go one")
		wb.Done()
	}()
	go func() {
		mt.Lock()
		v := x
		v++
		x = v
		mt.Unlock()
		runtime.Gosched()
		fmt.Println("go two")
		wb.Done()

	}()
	go func() {
		mt.Lock()
		v := x
		v++
		x = v
		mt.Unlock()
		fmt.Println("go three")
		runtime.Gosched()
		wb.Done()

	}()
	wb.Wait()
	fmt.Println("main x", x)
	fmt.Println("Num CPU:", runtime.NumCPU())
	fmt.Println("Num goroutines", runtime.NumGoroutine())

}
