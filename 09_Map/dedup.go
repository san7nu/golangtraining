package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := make(map[string]bool)
	fmt.Println("Enter")
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		if s.Text() == "end" {
			break
		}
		if lines[s.Text()] == false {
			lines[s.Text()] = true
		}
	}

	fmt.Println("Duplicate lines removed")
	for k, _ := range lines {
		fmt.Println(k)
	}

}
