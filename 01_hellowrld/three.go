package main

import "fmt"

func main() {

	a := make([]int, 0)
	a = append(a, 21)

	fmt.Println(a)

	addval(&a)

	fmt.Println(a)

}

func addval(b *[]int) {
	*b = append(*b, 35)
	return
}
