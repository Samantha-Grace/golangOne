package main

import "fmt"

func main() {
	// name := [count]Type{}
	scores := [5]int{15, 18, 17, 20, 30} // Declare and initiate in one call
	// scores[0] -> 15
	// scores[1] -> 18
	// scores[2] -> 17
	// scores[3] -> 20
	// scores[4] -> 30 (length of array - 1)

	//fmt.Println("First item:", scores[0])
	//fmt.Println("Last item:", scores[4])

	//fmt.Println("count of items:", len(scores))
	//fmt.Println("capacity of items:", cap(scores))

	iterate(scores)

	secondArr := make([]int, 5) // Declaring the array
	// secondArr[0] = 0
	// secondArr[1] = 0
	// secondArr[2] = 0
	secondArr = []int{5, 25, 80} // Initiating the array
	fmt.Println("count of items:", len(secondArr))
	fmt.Println("capacity of items:", cap(secondArr))

	scores[2] = 250
	iterate(scores)
}

// takes an array with 5 length
func iterate(myArr [5]int) {
	for i := 0; i < len(myArr); i++ {
		fmt.Println(i, myArr[i])
	}
	/*
		i  myArr[i]
		0 15
		1 18
		2 17
		3 20
		4 30
	*/
}
