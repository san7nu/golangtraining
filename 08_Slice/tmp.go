package main

import "fmt"

func main() {
	a := []int{1, 2, 4, 3, 6, 9}
	i := 0

	for i, _ := range a {
		fmt.Print(i)	
	}
	fmt.Println("Post-range1", i)

	for i = 0; i < len(a); i++ {
	}
	fmt.Println("Post-for", i)

	for i, _ = range a {
	}
	fmt.Println("Post-range2", i)
}
