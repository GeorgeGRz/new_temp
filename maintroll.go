package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Printf("%s,%s", runtime.GOOS, runtime.GOARCH)
}
