package main

import "fmt"

// We can pass exactly 1 value per parameter
// rand.Intn(100)
// We can send more than 1 ( in slices and arrays)
// Sen an array or slice entirely into a function

// When we want to send 0 or more values  to a function
//

// p ...string
// This means we can send no string or any count of string to the function.
// like sending nothing or sending a slice to the function
func display(p ...string) {
	fmt.Println("count of variables:", len(p))
	for _, v := range p {
		fmt.Println(v)
	}
}

func displayV2(a int, b float64, x bool, p ...string) {
	fmt.Println("count of variables:", len(p))
	for _, v := range p {
		fmt.Println(v)
	}
}

func main() {

	display()

	display("Samantha")

	display("Samantha", "Nima")

	display("one", "two", "three")

	names := []string{
		"1st",
		"2nd",
		"3rd",
		"4th",
	}

	// names... means like this: "1st",	"2nd","3rd","4th"
	display(names...) // Here is the most important point! use ... if you want to send a slice!!!!!

	s1 := []int{55}
	s2 := []int{2, 1, 8, 4, 6}

	s1 = append(s1, 78)

	s1 = append(s1, s2...)
	// s1 = append(s1, 2,18,4,6)
	//  s2... this means := 2,1,8,4,6

	fmt.Println(s1)
}
