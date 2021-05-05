package main

import (
	"fmt"
	"strconv"
)

func main() {
	//str := strconv.Itoa(12) // "12"

	intVal, err := strconv.Atoi("1A2")
	if err != nil { // there was an error
		panic(err) // Prints the error and exits
	}

	fmt.Println(intVal)
}
