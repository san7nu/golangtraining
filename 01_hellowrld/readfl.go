package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("dup.go")
	en := bufio.NewScanner(f)

	for en.Scan() {
		fmt.Println(en.Text())
	}

}
