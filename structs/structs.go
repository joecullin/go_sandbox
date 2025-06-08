package main

import "fmt"

type Vertex struct {
    X, Y int
    Z int
}

func main() {
    v := Vertex{1,4,0}
    v.Y = 7
    fmt.Println(v)

    p := &v
    p.X = 3

    fmt.Println(v)

    q := Vertex{X: 0}
    fmt.Println(q)
}
