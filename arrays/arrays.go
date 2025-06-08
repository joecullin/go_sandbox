package main

import (
	"fmt"
	"strings"
)

func inner() {
	// var a [2]string
	a := [2]string{"", ""}
	a[0] = "hey"
	fmt.Println(a)

	b := [7]int{0, 0, 1, 2, 4, 8, 300}
	b[0] = 1
	fmt.Println(b)

	c := b[0:3]
	c[0] = 500
	fmt.Println(b)

	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{1, true},
		{1, true},
		{1, true},
		{2, false},
		{3, false},
		{4, false},
	}

	fmt.Println(s)

	s = s[:0]
	fmt.Println(s)

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered from panic: ", r)
	// 		fmt.Println("how'd ya like that panic??")
	// 	}
	// }()
	// s = s[:10]
	// fmt.Println(s)
}

func main() {
	inner()

	a := make([]int, 0, 10)
	fmt.Printf("%b len=%d cap=%d %v\n", a, len(a), cap(a), a)
	a = a[:10]
	a[3] = 4
	fmt.Printf("%b len=%d cap=%d %v\n", a, len(a), cap(a), a)

	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[1][1] = "X"
	for i := 0; i < len(board); i++ {
		row := board[i]
		// for j := 0; j < len(row); j++ {
		// 	fmt.Print(" ", row[j], " ")
		// }
		// fmt.Println()
		fmt.Println(strings.Join(row, " "))
	}

	// var s []int
	s := make([]int, 3, 30)
	printSlice("s", s)
	s3 := append(s, 50)
	printSlice("s", s)
	printSlice("s3", s3)
	s2 := append(s, 1, 2, 3)
	printSlice("s2", s2)
	printSlice("s", s)
}

func printSlice(label string, s []int) {
	fmt.Printf("%s: len=%d cap=%d %v\n", label, len(s), cap(s), s)
}
