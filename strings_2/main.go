package main

import (
	"fmt"
	"strings"
)

func main() {
	//name := "samantha is learning"
	name := "SAMANTHA IS LEARNING"
	// result := strings.ToTitle(name)
	// Samantha Is Learning
	result := strings.ToLower(name)
	fmt.Println("Original:", name)
	fmt.Println("Edited:", result)

	name1 := "Samantha"
	name2 := "samantha"

	is := strings.ToLower(name1) == strings.ToLower(name2) // approach1
	is = strings.EqualFold(name1, name2)
	fmt.Println("Result:", is)

	mainStr := "samantha is learning golang"
	subStr := "is learning"
	res := strings.Contains(mainStr, subStr)
	fmt.Println("contains:", res)
}
