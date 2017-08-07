package main

import "fmt"

func main() {

	a := []int{8, 4, 5, 3, 1, 9}
	var b []int
	var posi, tmp int

	fmt.Printf("%T \n", a)
	fmt.Println(len(a), a)

	for i := 0; i < len(a); i++ {
		tmp, posi = maxx(a)
		b = append(b, tmp)
		a[posi] = -1
	}
	fmt.Println(b)
}

func maxx(x []int) (int, int) {

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
