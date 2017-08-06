package main

import "fmt"

func main() {

	var a [10]int
	fmt.Printf("%T \n", a)

	for i := 0; i < 10; i++ {
		a[i] = i
	}

	fmt.Println("--- ", a)

}
