package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	//var mu sync.Mutex
	//counter := 0
	var counter2 int64
	const gs = 50

	wg.Add(gs)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter2, 1)
			/*
				mu.Lock()
				v := counter
				//runtime.Gosched()
				v++
				counter = v
				mu.Unlock()
					wg.Done()*/
			fmt.Println(atomic.LoadInt64(&counter2))
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("===================================")
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter2)

}
