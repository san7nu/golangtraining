package main

import "fmt"

func inc2() func() int {
	y := 0
	return func() int {
		y++
		return y
	}

}

func main() {

	x := 0

	inc := func() int {
		x++
		return x
	}

	funcinc := inc2()

	fmt.Println(inc())
	fmt.Println(funcinc())
}
