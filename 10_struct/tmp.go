package main

import "fmt"

func main() {
	type str1 struct {
		id int
	}
	type str2 struct {
		name string
	}
	type str3 struct {
		x str1
		y str2
	}

	a := str1{1}
	b := str2{"abc"}
	fmt.Println(a, b)

	var c str3
	c.x.id = 2
	c.y.name = "def"
	fmt.Println(c)
}
