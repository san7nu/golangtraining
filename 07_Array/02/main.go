package main

import "fmt"

func main() {

	b := [10]int{8, 5, 4, 1, 2, 6, 3}
	var a [10]int
	var posi int

	fmt.Printf("%T \n", b)

	for i := 0; i < len(b); i++ {
		a[i], posi = maxx(b)
		b[posi] = -1
	}
	fmt.Println(a)
}

func maxx(x [10]int) (int, int) {
	e := 0
	posi := 0
	for i, v := range x {
		if e < v {
			e = v
			posi = i
		}
	}
	return e, posi
}
