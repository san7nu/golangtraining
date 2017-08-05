package main

import "fmt"

func main() {

	fmt.Println(abc(1, 2, 3))
}

func abc(n ...int) int {

	sum := 0

	for _, x := range n {
		sum += x
	}
	return sum
}
