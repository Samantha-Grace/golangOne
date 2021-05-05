package main

import "fmt"

func main() {
	for i := 0; i < 10; i += 2 {
		fmt.Println(i)
	}

	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

	for {
		fmt.Println("Help !!!")
	}

}
