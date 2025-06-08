package main

import "fmt"
import "time"

type MyError struct {
    When time.Time
    What string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s", e.When, e.What)
} 

func run() (int, error) {
    counter := 1
    if counter % 2 == 0 {
        return counter, nil
    } else {
        return 0, &MyError{
            time.Now(),
            "uh-oh!!!!",
        }
    }
}

func main() {
    if counter, err := run(); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("it worked!", counter)
    }
}
