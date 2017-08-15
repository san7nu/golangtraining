package main

import "fmt"

func main() {

	const a = 1
	const b = '1'
	const c = "1"
	// Constant can only b int, rune, string, boolean

	fmt.Printf("%T %T %T\n", a, b, c)

}
