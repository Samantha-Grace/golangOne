// A- Declare an empty array with any length, then fill that array by user inputs. ( Read from Scanf() )

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// import "fmt"

/*func main() {
	fmt.Println(cats("Jeremy", 9, 2, 4.5, "dinosaur", true))
}

func cats(name string, age int, sisters int, food float64, toy string, boy bool) string {
	result :=
		fmt.Sprintf("%s is %v years old and he has %v sisters. He eats %.1fkg of food a day and his favourite toy is a %s. Is he a good boy? %t!",
		name, age, sisters, food, toy, true)

	return result
}*/

// B- Assume we have an array with random numbers in a range of 1 to 100 and with a length of 100.
// Iterate over that and print the even numbers.

// import (
// 	"fmt"
// 	"math/rand"
// )

func main() {
	rand.Seed(time.Now().Unix()) //init the random.

	var arr = [100]int{} //declare and initiate array.

	arr = fillArray(arr) //by value
	//fillArrayV2(&arr) //by reference

	fmt.Println("Maximum amount is:", max(arr))
	fmt.Println("Sum is:", sum(arr))
}

/*func main() {
	rev()
}*/
func fillArray(a [100]int) [100]int {
	fmt.Printf("\n type: %T", a)
	for i := 0; i < 100; i++ {
		a[i] = rand.Intn(100) // i-th index of the array
	}

	return a
}

func fillArrayV2(arr *[100]int) {

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) // i-th index of the array
	}
}

//C- Assume we have two arrays with a length of 10. One of them is empty and another one is initialized.
//Iterate over full one and fill the empty one with  items of full  one, but in reverse order.
//Example:
//Full : 1,2,3,4,5
//Empty must be : 5,4,3,2,1
//
//package main

// import "fmt"

func rev() {

	arrOne := [10]int{}
	arrTwo := [10]int{10, 11, 12, 13, 14, 15, 66, 97, 78, 19}
	fmt.Println(arrOne, arrTwo)

	//arrOne = arrTwo
	j := 0
	for i := len(arrTwo) - 1; i >= 0; i-- {
		//fmt.Println(arrOne[i])
		// i = 9 ;arrTwo[i]=19
		//arrOne[0] = arrTwo[9]
		//arrOne[1] = arrTwo[8]
		//arrOne[2] = arrTwo[7]
		//arrOne[3] = arrTwo[6]
		//arrOne[4] = arrTwo[5]
		//arrOne[5] = arrTwo[4]
		arrOne[j] = arrTwo[i]
		j++
	}

	fmt.Println(arrOne, arrTwo)
}

// Array
// a[0] , a[1] a[n]
// 0 ,1 , n (indices)
// a[0] = 12 , a[1] = 45 , a[n] = 100
// Values

func sum(a [100]int) int {
	var sum = 0

	for _, v := range a {
		sum += v
	}

	return sum
}

func max(a [100]int) int {
	max := a[0]

	for _, v := range a {
		if v > max {
			max = v
		}
	}

	return max
}
