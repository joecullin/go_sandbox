package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	words := strings.Split(s, " ")
	fmt.Println("words", words)

	wordCounts := make(map[string]int)
	for _, word := range words {
		fmt.Println("word:", word)
		wordCounts[word]++
	}

	return wordCounts
}

func main() {
	fmt.Println("result", WordCount("one two three"))
	fmt.Println("result", WordCount("a gopher's gotta do what a gopher's gotta do"))
}