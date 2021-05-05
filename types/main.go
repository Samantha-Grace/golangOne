package main

import "fmt"

func main() {
	var (
		i  = 12
		f  = 12.25
		z  = "Samantha"
		ok = true
	)

	fmt.Printf("\n i Type is : %T and its value is %v", i, i)

	fmt.Printf("\n f Type is : %T and its value is %v", f, f)

	fmt.Printf("\n z Type is : %T and its value is %v", z, z)

	fmt.Printf("\n ok Type is : %T and its value is %v", ok, ok)

}
