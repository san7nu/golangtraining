package main

import "fmt"

func main() {
	s := "ths is str"
	fmt.Println(s)
	fmt.Println(s[0])
	s[0] = 'B'
	fmt.Println(s[0])
}
