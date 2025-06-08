package main


import (
    "fmt"
    "math"
    // "errors"
)

type ErrNegativeSqrt float64

func (badValue ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("Error! can't get square root of negative %v", fmt.Sprint(float64(badValue)))
}

func Sqrt(x float64) (float64, error) {
    if (x < 0){
        // return 0, errors.New(fmt.Sprintf("Error! can't get square root of negative %v", x))
        return 0, ErrNegativeSqrt(x)
    }

    guess := 1.0
    prevGuess := 1.0
    for i:=0; i<10; i++ {
        guess -= (guess*guess - x) / (2*guess)
        fmt.Printf("i=%d guess: %32.30f\n", i, guess)
        if math.Abs(prevGuess - guess) < 0.000000000001 {
            break
        }
        prevGuess = guess
    }
    return guess, nil
}

func main() {
    for i:=-1; i<=3; i++ {
        sqrt, err := Sqrt(float64(i))
        if (err != nil){
            fmt.Println(i, err)
        } else {
            fmt.Println(i, sqrt)
        }
    }
}
