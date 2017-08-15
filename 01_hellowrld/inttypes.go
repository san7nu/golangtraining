package main

import "fmt"

func main() {
	var a int8
	var b uint8
	a, b = 127, 127
	fmt.Println(a, b)
	a++
	b++
	fmt.Println(a, b)
}
