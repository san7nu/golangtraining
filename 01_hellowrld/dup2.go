package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]int)
	en := bufio.NewScanner(os.Stdin)

	if en.Scan() {
		m[en.Text()]++
	}

	for i, v := range m {
		fmt.Println(i, v)
	}
}
