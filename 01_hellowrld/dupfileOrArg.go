package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args[1:]) > 1 {
		fmt.Println(os.Args[1])
		f, _ := os.Open(os.Args[1])
		countlines(f)
	} else {
		fmt.Println("INSIDE ELSE")
		countlines(os.Stdin)
	}

}

func countlines(f *os.File) {
	out := make(map[string]int)
	en := bufio.NewScanner(f)

	for en.Scan() {
		out[en.Text()]++
		if en.Text() == "end" {
			break
		}
	}

	for k, v := range out {
		fmt.Println(k, v)
	}
}
