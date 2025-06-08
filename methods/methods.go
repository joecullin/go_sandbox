package main

import "fmt"
import "math"

type Vertex struct {
    X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("before", Abs(v))
	Scale(v, 10)
	fmt.Println(Abs(v))
}


// func (v *Vertex) Scale(factor float64) {
//     fmt.Println("called scale with factor", factor)
//     v.X = v.X * factor
//     v.Y = v.Y * factor
// }
// 
// func main() {
//     v := Vertex{3,4}
//     fmt.Println("initial:", v)
// 
//     v.Scale(2)
//     fmt.Println("scaled:", v)
// }
