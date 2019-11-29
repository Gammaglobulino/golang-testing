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
		runtime.Gosched()
		fmt.Println("go one")
		x = v
		wb.Done()
		mt.Unlock()
	}()
	go func() {
		mt.Lock()
		v := x
		v++
		runtime.Gosched()
		fmt.Println("go two")
		x = v
		wb.Done()
		mt.Unlock()
	}()
	go func() {
		mt.Lock()
		v := x
		v++
		fmt.Println("go three")
		runtime.Gosched()
		x = v
		wb.Done()
		mt.Unlock()
	}()
	wb.Wait()
	fmt.Println("main x", x)
	fmt.Println("Num CPU:", runtime.NumCPU())
	fmt.Println("Num goroutines", runtime.NumGoroutine())

}
