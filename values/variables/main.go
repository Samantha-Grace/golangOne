package main

import "fmt"

/*func main() {

	w := 12 // 12.000
	fmt.Println("W=", w)

	// Pointer
	// is a variable that keeps the memory address of another variable
	var pw2 *int
	*pw2 = 100

	pw := &w
	fmt.Println(pw)

	*pw = 62
	fmt.Println("W=", w)
}*/

func main() {
	//displayMe()
	//sam("Hello, Samantha!")
	//sam("I am learning Golang")
	//msg := sam2("Samantha")
	//fmt.Println(msg)

	fmt.Println(sam3(3, 7))

	v := sam3(30, 70)
	fmt.Println(v)
}

func displayMe() {
	fmt.Println("I'm here!")
}

// name => sam
// Parameters => greeting
func sam(greeting string) {
	fmt.Println(greeting)
}

// name
// Parameters
// Return value
func sam2(name string) string {
	return "Hello," + name
}

func sam3(x int, y int) int {
	return x * y
}
