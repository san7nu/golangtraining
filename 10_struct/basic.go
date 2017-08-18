package main

import "fmt"

func main() {
	type stu struct {
		id int
		name string
	}

	type game struct {
		stu
		status bool
	}

	x := stu{1, "aaa"}
	var y stu
	y.id = 2
	y.name = "bbb"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x == y)
	
	g1 := game{stu{2,"bbb"},true}
	var g2 game
	g2.id = 3
	g2.name = "ccc"
	g2.status = false 
	fmt.Println(g1)
	fmt.Println(g2)
}
