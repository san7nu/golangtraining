package main

import "fmt"

func main() {

	//var u, v string

		func abc(a, b string) (x, y string) {
		x = a
		y = b
		return
	}


	u, v := abc("1", "2")

	fmt.Println(u, v)

}

