package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	go foo(c)

	wg.Add(1)
	go bar(c)
	wg.Wait()
	fmt.Println("exits")
}
func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println(<-c)
	wg.Done()
}
