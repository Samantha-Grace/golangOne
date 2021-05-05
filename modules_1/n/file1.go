package n

import "fmt"

var X int = 100
var x int = 200

// Func1
// Exported ...
func Func1() {
	fmt.Println("I am func 1 from package n")
	fmt.Println("Y=", X)
	func1()
}

func func1() {

}
