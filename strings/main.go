package main

import (
	"fmt"
)

func main() {
	i := 2
	//strVal := strconv.Itoa(i) // Int to ASCII code

	//fmt.Println(strVal)

	a, b := sqrtt(i)
	fmt.Printf("\n num^2 = %d , num^3= %d", a, b)
}

func sqrtt(num int) (int, int) { // Return values with name
	i := num * num       // num^2
	j := num * num * num // num^3

	return i, j
}
