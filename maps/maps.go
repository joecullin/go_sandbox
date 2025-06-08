package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex


func main() {
    // m = make(map[string]Vertex)
    m := map[string]Vertex {
        "Somewhere, NJ": Vertex{25.3, -74.39967},
        "Here, PA": {25.3, -74.39967},
    }

    m["Bell labs"] = Vertex{40.68433, -74.39967}
    m["test joe"] = Vertex{1.0,0.0}

    fmt.Println(m)
    fmt.Println(m["Bell labs"])

    if elem, ok := m["joe"]; ok {
        fmt.Println("found joe!", elem)
    }else {
        fmt.Println("no joe!")
    }
}
