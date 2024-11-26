package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	for i := 65; i <= 90; i++ {
		b.WriteByte(byte(i))
	}
	runes := []rune(b.String())
	for i, r := range runes {
		if i == 14 {
			fmt.Printf("Привет")
			continue
		}
		if (i+1)%3 == 0 {
			fmt.Printf("%d", i+1)
			continue
		}
		if (i+1)%5 == 0 {
			index := i + 1040
			r = rune(index)
		}
		fmt.Printf("%c", r)
	}
}
