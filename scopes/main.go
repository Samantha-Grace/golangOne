package main

import "fmt"

const c1 = "hello" // type inference

const (
	c2 = (iota + 1) * 10 // (0+1) *10
	c3                   // 20
	c4                   // 30
)

func main() {
	var (
		x int = 10
		y int = 20
	)

	fmt.Println(c2, c3, c4)
	// new , is a keyword that declares a pointer to the provided type
	// a:= new(int)

	fmt.Printf("\n x = %d , y=%d", x, y)

	// swap(x, y)
	swapByReference(&x, &y)

	fmt.Printf("\n x = %d , y=%d", x, y)
}

// Call by Value
func swap(x int, y int) {
	temp := y
	y = x
	x = temp

	fmt.Printf("\n x = %d ,y=%d", x, y)
}

// Call by Reference.
func swapByReference(x *int, y *int) {
	temp := y
	y = x
	x = temp

	fmt.Printf("\n x = %d ,y=%d", *x, *y)
}
