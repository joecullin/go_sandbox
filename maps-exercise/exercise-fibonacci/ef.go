package main

import "fmt"

// func next() func() int {
// 	current := [2]int{0,1}

// 	return func () int {
// 		nextVal := current[0] + current[1]
// 		current[0] = current[1]
// 		current[1] = nextVal
// 		return current[1]
// 	}
// }

func next() func() int {
	f1, f2 := 0, 1

	return func () int {
		returnVal := f1
		nextVal := f1 + f2
		f1 = f2
		f2 = nextVal
		return returnVal
	}
}

func main() {
	seq := next()
	for range 20 {
		fmt.Println(seq())
	}
}