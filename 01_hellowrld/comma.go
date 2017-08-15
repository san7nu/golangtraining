package main

import (
	"fmt"
)

func main() {
	str := "0123456789"
	out := make([]string, 0)

	for i := 0; i < len(str); i += 3 {
		if i+3 > len(str) {
			out = append(out, str[i:])
		} else {
			out = append(out, str[i:i+3]+",")
		}
	}
	fmt.Println(out)
}
