package main

import (
	"fmt"
	"github.com/GeorgeGRz/new_temp/cat"
	"runtime"
)

func main() {

	fmt.Printf("%s,%s\n", runtime.GOOS, runtime.GOARCH)
	cat.Somefunc()

}
