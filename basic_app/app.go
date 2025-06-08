package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	prog := os.Args[0]
	args := os.Args[1:]

	showHelp := false
	flag.BoolVar(&showHelp, "h", false, "show help")

	showVersion := false
	flag.BoolVar(&showVersion, "version", false, "show version")

	var flagN = flag.Int("n", 0, "number of things, i.e. how many.")
	var flagDirection = flag.String("d", "up", "direction: [up|down]. default is up.")

	flag.Parse()

	fmt.Println("prog:", prog)
	fmt.Println("params:", args)
	fmt.Println("showHelp:", showHelp)
	fmt.Println("showVersion:", showVersion)
	fmt.Println("n:", *flagN)
	fmt.Println("direction:", *flagDirection)
}
