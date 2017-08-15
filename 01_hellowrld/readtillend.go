package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	n := bufio.NewScanner(os.Stdin)
	s := make([]string,0)

	for n.Scan() {
		if n.Text() == "end" {
			break
		} else {
			s = append(s, n.Text())
		}
	}
	fmt.Println(s)
}
