package main

import "fmt"

func main() {
	n := "Samantha"

	greeting(myName, n)

	greeting(sayBye, n)
}

func greeting(n func(string), name string) {
	n(name)
}

func myName(name string) { // Type of myName is : func(string)
	fmt.Println("Hello,", name)
}

func sayBye(name string) {
	fmt.Println("Bye ", name)
}
