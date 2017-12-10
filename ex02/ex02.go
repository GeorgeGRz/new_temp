package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println(p.first, "says hello!")
}

type human interface {
	speak()
}

func saySomething(a human) {
	a.speak()
}
func main() {
	p1 := person{"George", "Moustakas", 17}
	p1.speak()
	saySomething(&p1)
}
