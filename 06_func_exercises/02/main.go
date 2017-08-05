package main

import "fmt"

func main() {
	var in int
	fmt.Print("-----> ")
	fmt.Scan(&in)
	abc := func(x int) (n int, b bool) {
		return x / 2, x%2 == 0
	}

	fmt.Println(abc(in))
}
