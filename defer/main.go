package main

import "fmt"

func main() {
	fmt.Println("1-main") // First
	fmt.Println("2-main")

	// Stack
	// LIFO (Last In First Out)
	defer f1()
	defer f2()

	// Anonymous functions
	// Or function closure.
	defer func() {
		fmt.Println("I am an anonymous function that deferred!!!!!")
	}()

	fmt.Println("3-main")
	fmt.Println("4-main")
	fmt.Println("5-main") // Last
}

func f1() int {
	defer fmt.Println("1-f1 just before return") // last
	fmt.Println("1-f1")                          // First
	fmt.Println("2-f1")
	fmt.Println("3-f1")
	fmt.Println("4-f1")

	return 1212121
}

func f2() int {
	defer fmt.Println("1-f2 just before return") // last
	fmt.Println("1-f2")                          // First
	fmt.Println("2-f2")
	fmt.Println("3-f2")
	fmt.Println("4-f2")

	return 1212121
}
