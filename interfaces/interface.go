package main

import "fmt"

type Person struct {
    Name string
    Age int
}

// A "Stringer" interface
func (p Person) String() string {
    return fmt.Sprintf("%v: %v years; ", p.Name, p.Age);
}

func main() {
    joe := Person{"joe", 48}
    barney := Person{"Barney", 9}
    fmt.Println(joe, barney)
}
