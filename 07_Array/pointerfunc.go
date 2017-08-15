package main

import "fmt"

func main() {
	var b *string
	a := "abcd"
	b = &a
	fmt.Println("b =", b)
	fmt.Println("*b =", *b)
	whtEz(b)
	fmt.Println("b =", b)
	fmt.Println("*b =", *b)
}

func whtEz(n *string) {
	fmt.Println("INSIDE")
	fmt.Println("n =", n)
	fmt.Println("*n =", *n)
	*n = "20"
	fmt.Println("*n =", *n)
}
