package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
func foo() {
	fmt.Println("Foo says hello")
	wg.Done()
}
func bar() {
	fmt.Println("Bar says hello")
	wg.Done()
}
