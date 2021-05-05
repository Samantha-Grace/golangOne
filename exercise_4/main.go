// 4 - We have a map of scores of students of a course. The map is of names to float scores.
// Using this map;
//  declare a function to find and return the highest score and the name of that student.
//  declare a function to calculate and return the sum of scores.
// declare a function to calculate and return the average score of the course.

package main

import "fmt"

/*type Grade struct {
	Key   string
	Value float64
}

type GradeList []Grade

func (g GradeList) Len() int           { return len(g) }
func (g GradeList) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }
func (g GradeList) Less(i, j int) bool { return g[i].Value < g[j].Value }

func main() {

	scores := map[string]float64{
		"Able":   8.0,
		"Beth":   7.5,
		"Cain":   9.0,
		"Declyn": 10.0,
	}

	g := make(GradeList, len(scores))

	i := 0
	for k, v := range scores {
		g[i] = Grade{k, v}
		i++
	}

	sort.Sort(g)

	for _, k := range g {
		fmt.Printf("%v\t%v\n", k.Key, k.Value)
	}

	var total float64
	for _, v := range scores {
		total += v
	}
	fmt.Printf("The sum of the scores: %.02f\n", total)

	//Trying to figure out how to get key[3] to be Declyn's name in %s
	var high float64
	for k, v := range scores {
		high = v
		fmt.Printf("%s has the highest of the scores: %.02f\n", k, high)
	}
}*/

// 	var i int
// 	sorted := make(GradeList, len(scores))
// 	for key, value := range scores {
// 		sorted[i] = Grade{key, value}
// 		i++
// 		fmt.Printf("%s %2.1f\n", key, scores[key])
// 	}
// }

//sort maps by key
// 	keys := make([]string, 0, len(scores))

// 	for key := range scores {
// 		keys = append(keys, key)
// 	}

// 	sort.Strings(keys)

// 	for _, key := range keys {
// 		fmt.Printf("%s %2.1f\n", key, scores[key])
// 	}
// }

// 	type kv struct {
// 		Key   string
// 		Value float64
// 	}

// 	var grade []kv
// 	for k, v := range scores {
// 		grade = append(grade, kv{k, v})
// 	}

// 	sort.Slice(grade, func(i, j int) bool {
// 		return grade[i].Value > grade[j].Value
// 	})

// 	for _, kv := range grade {
// 		fmt.Printf("%s, %2.1f\n", kv.Key, kv.Value)
// 	}
// }

//-----------------------------------------
// func average(scores []float64) float64 {
// 	for _, v := range scores {
// 	}
// 	return float64(len(scores))
// }

// fmt.Println(scores)
// func highScore(scores float64) float64 {
// 	for i := 0; i <= 10.0; i++ {
// 	}
// 	return scores
// }

// func highScore(k string, v float64) float64 {
// 	for _, v := range k {
// 		fmt.Println(v)
// 	}
// 	return v
// }

// func highScore(k string, v float64) float64{
// if v
// }

// //add all v(values) together
// func sumScore() {

// }
// //add all v(values) and multiply by 4
// func averageScore() {

// }

// a := highScore(scores)
// fmt.Println(a)

// for i = 1; i < scores.length; i++ {
// 	if (scores.get(i) > max) {
// 	  max = scores.get(i);
// scores = math.Max(v)

// func highScore(k string, v float64) float64 {

// 	for maxNumber = range scores {
// 		break
// 	}
// 	for a := range scores {
// 		if a > maxNumber {
// 			maxNumber = a
// 		}
// 	}
// 	return maxNumber
// }

// for k, v := range scores {
// 	fmt.Printf("\n %s has a score of: %2.1f \n", k, v)
// }

// sort.Slice(scores,
// 	func(i, j int) bool {
// 		return scores[i] < scores[j]
// 	})
// sort.Slice(scores, sorter)

// fmt.Println(scores)

// }
// func sorter(i int, j int) bool {
// return scores[i] > scores[j]
// }

// func highScore(scores []float64) float64{
// 	for i := 0; i < 10.0 ; i++ {
// 		if scores[0] < scores[3] {
// 			fmt.Println(scores)
// 		}
// 	}
// }
// min, max := highScore(scores)

// func highScore(s []float64) float64 {
// 	max = scores[0]
// 	for _, score := range scores {
// 		if score.Grade > max.Grade {
// 			max = score
// 		}
// 	}
// 	return max
// }
// func highScore(names string, scores float64) float64 {
// 	for j := 1; j < temp; j++ {
// 		if scores[0] < scores[j] {
// 			scores[0] = scores[j]
// 		}
// 		return scores[0]
// 	}

// }

// size := 4
// sum :=0

// for k, v := range scores {
// 	fmt.Printf("%s has a score of: %2.1f \n", k, v)
// }
// //highScore(scores)
// }

// func sumScore(s []float64) float64 {
// 	for i := 0; i < size; i++ {
// 		sum += (scores[i])
//     }
// 	avg := (float64(sum)) / (float64(size))
// 	fmt.Println("Sum = ", sum, "\nAverage = ", avg)
// }

func main() {
	studentScores := map[string]float64{
		"Samantha": 10.0,
		"Somi":     9.0,
		"Jeni":     8.0,
		"Susan":    5.0,
		"Hannah":   15.0,
	}

	maxName, maxScore := highest(studentScores)
	fmt.Printf("Max Score is: %2.2f and the student name is: %s \n",
		maxScore, maxName)

	fmt.Println("The sum of all scores is :", sum(studentScores))

	ave := sum(studentScores) / float64(len(studentScores))

	fmt.Println("The average of all scores is :", ave)
}

func highest(scores map[string]float64) (string, float64) {
	maxScore := 0.0
	student := "John"

	for name, score := range scores {
		if score > maxScore {
			maxScore = score
			student = name
		}
	}

	return student, maxScore
}

func sum(scores map[string]float64) float64 {
	sumScores := 0.0

	for _, score := range scores {
		sumScores += score
		//sumScores = sumScores + score
	}

	return sumScores
}
