package main

import "fmt"

func main() {

	scores := map[string]float64{
		"listening": 6.7,
		"writing":   6.5,
		"speaking":  3.0,
		"reading":   7.0,
	}

	fmt.Println("my score of LIS is:", scores["listening"])
	fmt.Println("my score of talk is:", scores["talking"])

	for k, v := range scores {
		fmt.Printf("\n Skill is: %s and Score is: %2.1f \n", k, v)
	}

	for _, v := range scores {
		fmt.Printf("\n Score is: %2.1f \n", v)
	}

	val := scores["talking"]
	val, ok := scores["talking"]
	if !ok {
		fmt.Println("there is no value for <talking>")
		//_ = val
	} else {
		fmt.Println(val)
	}

	// Adding to a map
	// Is assigning a value to a key. or Updating the value of a key.
	scores["talking"] = 9.0

	val, ok = scores["talking"]
	if !ok {
		fmt.Println("there is no value for <talking>")
		//_ = val
	} else {
		fmt.Println("the value is added now:", val)
	}

	// Updating
	scores["listening"] = 10.0
	fmt.Println("listening is now:", scores["listening"])

	// Deleting from a map
	delete(scores, "talking")

	val = scores["talking"]
	val, ok = scores["talking"]
	if !ok {
		fmt.Println("there is no value for <talking>")
		//_ = val
	} else {
		fmt.Println("the value is added now:", val)
	}

	val, ok = scores["talking"]
	if ok {
		delete(scores, "talking")
	} else {
		fmt.Println("There is no value for <talking>")
	}
	// Maps are not thread safe.
}
