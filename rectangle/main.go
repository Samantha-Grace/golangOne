package main

import "fmt"

func main() {
	w := 10
	l := 20

	//area(w, l)

	a := areaR(w, l)
	fmt.Println("area is:", a)
}

func area(width, length int) {
	fmt.Println("Area is :", length*width)
}

func areaR(w, l int) int {
	return w * l
}
