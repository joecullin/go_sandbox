package main

import (
    "fmt"
    "io"
)

type MyReader struct{}

func (r MyReader) Read(buffer []byte) (int, error) {
    // a := "A"
    // buffer = byte(a[0])
    buffer[0] = 'A'
    return 1, nil
}

func main() {

    var r MyReader
    b := make([]byte, 8)
    all := make([]string, 0, 0)
    for {
        n, err := r.Read(b)
        fmt.Printf("n = %v; err = %v; b = %v\n", n, err, b)
        fmt.Printf("b[:n] = %q\n", b[:n])
        if (err == io.EOF){
            break
        } else{
            all = append(all, "Z")
            // all := append(all, string(b[0]))
            fmt.Printf("all (%v) = %v\n", len(all), all)
        }
    }
}


// 
//     r := strings.NewReader("Hello, Reader!")
//     
//     b := make([]byte, 8)
//     for {
//         n, err := r.Read(b)
//         fmt.Printf("n = %v; err = %v; b = %v\n", n, err, b)
//         fmt.Printf("b[:n] = %q\n", b[:n])
//         if (err == io.EOF){
//             break
//         }
//     }
// }
