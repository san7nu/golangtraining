package main

import "fmt"

func main() {

	i := []int{15, 62, 86, 55, 4, 99, 66}
	fmt.Println(grt(i...))

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
