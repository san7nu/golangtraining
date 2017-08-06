package main

import "fmt"

func main() {

	b := [10]int{8, 5, 4, 1, 2, 6, 3}

	fmt.Printf("%T \n", b)

	for i := 0; i < len(b) && b[i]!=0 ; i++ {
		fmt.Println("----> ", b[i])
	}

}
