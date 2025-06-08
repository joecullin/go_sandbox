package main

import (
    "fmt"
//    "math"
)

func main() {
    sum := 0
    i := 0
    // for i:=0; i<10; i++ {
    for sum<10 {
        i++
        sum += i
        fmt.Println("sum", sum)
    }
    fmt.Println("final sum", sum)

    j := 0
    for {
        j++
        fmt.Println("aaaah", j);
        if j > 10 {
            break
        }
    }

    // vars inside an if/else. Like perl's "my". I do miss that, but not going to go deep on it now.
}





//    var sleepy, excited, nervous, hungry  bool
//    
//    func initBools() {
//        sleepy = true
//        excited = true
//        nervous = true
//    }
//    
//    func wakeUp() {
//        sleepy = false
//    }
//    
//    func add(x, y int) int {
//        return x + y;
//    }
//    
//    func swap (a, b string) (string, string) {
//        return b, a;
//    }
//    
//    // naked return
//    func double (a int) (x, y int) {
//        x = a
//        y = a
//        return
//    }
//    
//    
//    func main() {
//        fmt.Println("heeeyyy!!")
//        fmt.Println("pi:");
//        fmt.Println(math.Pi)
//        fmt.Println("that's it!")
//        fmt.Println("added:")
//        fmt.Println(add(1,2))
//    
//        a, b := swap("first", "second")
//        fmt.Println(a, b)
//    
//        fmt.Println("double")
//        fmt.Println(double(3))
//    
//        initBools()
//        fmt.Println("vars:")
//        fmt.Println(sleepy, excited, nervous, hungry)
//        wakeUp()
//        fmt.Println("are you awake now??:")
//        fmt.Println(sleepy, excited, nervous, hungry)
//    
//        // var f float, i int, b bool, s string
//        var f float64
//        var i int
//        var b2 bool
//        var s string
//        fmt.Printf("%v %v %v %q\n", i, f, b2, s)
//    
//        v := 42.0
//        fmt.Printf("v type is %T\n", v)
//        var v2 = int(v)
//        fmt.Printf("v2 type is %T\n", v2)
//    
//        const Yes = true
//        fmt.Println("Yes???", Yes) 
//    }
