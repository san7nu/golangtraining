package main

import "fmt"

func main() {
	var x int
	fmt.Print("----->")
	fmt.Scan(&x)
	fmt.Println(abc(x))
}

func abc(x int) (n int, b bool) {
	return x / 2, x%2 == 0
}
