package main

import (
	"mod1/m"
	"mod1/n"
)

func main() {
	n.X = 1
	n.Func1()

	m.Y = 5
	m.Func2()
}
