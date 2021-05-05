package main

import "fmt"

func main() {

	fmt.Println("Enter your name:")
	var name string
	fmt.Scanln(&name)
	fmt.Println(twoferV2(name))
}

func twofer(name string) string {

	temp := "you"
	if name != "" {
		temp = name
	}

	return "One for " + temp + ", one for me."
}

func twoferV2(name string) string {

	temp := "you"
	if name != "" {
		temp = name
	}

	result := fmt.Sprintf("One For %s One for me", temp)

	return result
}
