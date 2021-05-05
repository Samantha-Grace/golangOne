// A- Declare an empty array with any length, then fill that array by user inputs. ( Read from Scanf() )

// package main

// import "fmt"

// func main() {
// 	fmt.Println(cats("Jeremy", 9, 2, 4.5, "dinosaur", true))

// }

// func cats(name string, age int, sisters int, food float64, toy string, boy bool) string {
// 	result :=
// 	fmt.Sprintf("%s is %v years old and he has %v sisters. He eats %.1fkg of food a day and his favourite toy is a %s. Is he a good boy? %t!", name, age, sisters, food, toy, true)

// 	return result
// }

// B- Assume we have an array with random numbers in a range of 1 to 100 and with a length of 100.
// Iterate over that and print the even numbers.

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	rand.Seed(time.Now().Unix())
// 	var arr = [100]int{}
// 	arr = fillArr(arr)
// 	//fmt.Println(arr)
// 	for _, v := range arr {
// 		if v%2 == 0 {
// 			fmt.Println(v)
// 		}
// 	}
// }

//----

// func fillArr(arr [100]int) [100]int {
// 	for i := 0; i < 100; i++ {
// 		arr[i] = rand.Intn(100)
// 	}
// 	return arr
// }

// func fillArr2(arr [100]int) [100]int {

// 	for _,v := range arr{
// v=rand.Intn(100)
// print(v)
// 	}

// 	arr[i] = rand.Intn(100)

// 	return arr
// }

// C- Assume we have two arrays with a length of 10. One of them is empty and another one is initialized.
// Iterate over full one and fill the empty one with  items of full  one, but in reverse order.
// Example:
// Full : 1,2,3,4,5
// Empty must be : 5,4,3,2,1

package main

import "fmt"

func main() {

	arrOne := [10]int{}
	arrTwo := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arrOne, arrTwo)

	//arrOne = arrTwo
	for i := len(arrTwo) - 1; i >= 0; i-- {
		//fmt.Println(arrOne[i])

	}
}
