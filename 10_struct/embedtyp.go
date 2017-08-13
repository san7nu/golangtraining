package main

import "fmt"

type model struct {
	string
	id int
}

func main() {
	a := model{"2014", 133}
	fmt.Println(a)

// Cannot range over struct like below
//	for i, v := range a {
//		fmt.Println(i, v)
//	}
}
