package main

import "fmt"

func main() {

	fmt.Println(grt(67, 37, 89, 66, 53, 92, 18))
}

func grt(n ...int) int {
	tmp := 0
	for _, x := range n {
		if x > tmp {
			tmp = x
		}
	}
	return tmp
}
