package main

import "fmt"

func main() {
	var a string
	fmt.Println("enter your fav lang:")
	fmt.Scanln(&a)

	switch a {
	case "c":
		fmt.Println("your fav language is C!")
	case "go":
		fmt.Println("your fav language is Go!")
	case "java":
		fmt.Println("your fav language is Java!")
	default:
		fmt.Println("what is your fav language!???")
	}
}
