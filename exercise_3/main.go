// 3 - Assume a slice. All of the items of it are repeated twice except 1 item that is alone.
// Easy challenge: Find that number and print it.
// Hard challenge: Find that number using just 1 iteration on the slice.
// For example:
// a=[1,5,2,5,100,2,85,1,100]
// The exception is: 85

package main

import "fmt"

func main() {
	var ar = []int{20, 20, 3, 5, 4, 11, 8, 11, 5, 3, 4}
	fmt.Println(findSingle(ar))

}

// Finding the exception number using the XOR.
func findSingle(ar []int) int {
	res := ar[0]
	for i := 1; i < len(ar); i++ { // Started from 1  because we used the 0-index in above line.
		res = res ^ ar[i]
	}

	return res
}
