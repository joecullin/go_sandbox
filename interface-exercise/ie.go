package main

import "fmt"
import "strings"

type IPAddr [4]byte

func (ip IPAddr) String() string {
    var ipTokens = make([]string, len(ip));
    for i := range len(ip) {
        ipTokens[i] = fmt.Sprint(int(ip[i]))
    }
    return strings.Join(ipTokens, ".")
}

func main() {
    hosts := map[string]IPAddr {
        "loopback": {127, 0, 0, 1},
        "dns": {8, 8, 8, 8},
    }

    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
