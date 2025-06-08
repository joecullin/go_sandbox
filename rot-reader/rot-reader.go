package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(buffer []byte) (int, error) {
    n, err := r.r.Read(buffer)
	if err != nil {
        return 0, err
	}
    for i := range buffer {
        val := buffer[i]
        if (val > 'A' && val <= 'M') || (val > 'a' && val <= 'm'){
            buffer[i] += 13
        } else if (val >= 'N' && val <= 'Z') || (val >= 'n' && val <= 'z'){
            buffer[i] -= 13
        }
    }
    return n, nil
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
