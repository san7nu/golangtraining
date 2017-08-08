package main

import "fmt"

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	incA := wrapper()
	incB := wrapper()
	fmt.Println("A:", incA())
	fmt.Println("A:", incA())
	fmt.Println("B:", incB())
	fmt.Println("B:", incB())
	fmt.Println("B:", incB())
}
