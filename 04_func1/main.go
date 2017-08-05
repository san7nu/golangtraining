package main

import "fmt"

func main() {
	fmt.Println(abc(1))
}

func abc(x int) (n int, b bool) {
	return x / 2, x%2 == 0
}
