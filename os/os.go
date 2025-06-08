package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Current OS:")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("darwin (macos)")
	default:
		fmt.Printf("%s\n", os)
	}

}
