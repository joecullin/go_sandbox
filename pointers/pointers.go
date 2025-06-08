package main

import "fmt"

func main() {
    i, j := 42, 2701

    p := &i
    fmt.Println("p", *p)
    *p = 18
    fmt.Println("p", *p)

    q := &j
    fmt.Println("q", *q)
    // *q = *p
    q = p
    fmt.Println("q", *q)
}
