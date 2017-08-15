package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	if len(os.Args[1:]) >= 1 {
		fmt.Println(os.Args[1])
		f, _ := os.Open(os.Args[1])
		fmt.Println("FILENAME:",f.Name())
		countlines(f)
		f.Close()
	} else {
		fmt.Println("INSIDE ELSE")
		countlines(os.Stdin)
	}

}

func countlines(f *os.File) {
	out := make(map[string]int)
	en := bufio.NewScanner(f)

	for en.Scan() {
		if en.Text() == "end" {
			break
		}
		out[en.Text()]++
	}

	for k, v := range out {
		if v > 1 {
			fmt.Println(k, v)
		}
	}
}
