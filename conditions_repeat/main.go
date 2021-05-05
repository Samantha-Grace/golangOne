package main

import "fmt"

func main() {
	var c int
	fmt.Println("enter a number:")
	fmt.Scanf("%d", &c)

	if c < 10 {
		fmt.Println("Number is little than 10")
	} else if c < 100 {
		fmt.Println("Number is greater than 10 and little than 100")
	} else {
		fmt.Println("Number is greater than 10 and greater than 100")
	}

	/*	for i := 0; i < 100; i += 2 {
		fmt.Println(i)

	}*/
}
