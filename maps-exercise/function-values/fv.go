package main

import "fmt"
import "math"

func compute (fn func(float64, float64) float64) float64 {
	return fn(3,4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i :=0; i<5; i++ {
		fmt.Println(
			pos(1),
			// pos(i+2),
			neg(i * -1),
		)
	}
}