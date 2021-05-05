package main

import (
	"fmt"
	"sort"
)

var s2 = make([]int, 10)

func main() {
	s1 := []int{}

	fmt.Println("s1 length:", len(s1))
	fmt.Println("s1 cap:", cap(s1))
	fmt.Println("s2 length:", len(s2))
	fmt.Println("s2 cap:", cap(s2))
	fmt.Println("s2:", s2)

	// Appending
	s1 = append(s1, 85)
	fmt.Println("s1 length:", len(s1))
	fmt.Println("s1 cap:", cap(s1))
	fmt.Println(s1)

	s2 = append(s2, 85)
	s2 = append(s2, 99, 26)
	s2 = append(s2, 12, 200, 666, 2, 3, 4, 5, 6, 7, 2300, 120, 90)

	fmt.Println("s2 length:", len(s2))
	fmt.Println("s2 cap:", cap(s2))
	fmt.Println(s2)

	// a[i] , a[j]
	// a[i] < a[j]
	// s:=[]int{2,4,6,8,10} // Sorted Ascending
	sort.Slice(s2,
		func(i, j int) bool { // sequential indices of a slice
			return s2[i] < s2[j]
		})

	// Way 2, declare the function separately.
	sort.Slice(s2, sorter)

	fmt.Println(s2)
}

func sorter(i int, j int) bool {
	return s2[i] > s2[j] // multiplying
}

func displayMe(s []int) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// mySlice
// []int{0,0,0,0,0,0,0,0,0,0} => length = 10 and capacity = 10
//  Append , the appended element will be inserted at the end of the slice.
// Length of slice increases by 1 ( by count of appended elements)
// capacity of slice became double of previous capacity.

// init => []int{0,0,0,0,0,0,0,0,0,0}
// Append 85 => []int{0,0,0,0,0,0,0,0,0,0,85}
