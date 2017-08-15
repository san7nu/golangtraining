package main

import "fmt"

func main() {

	a := 15
	b := 10

	fmt.Printf("a      ::::: %b\n",a)
	fmt.Printf("b      ::::: %b\n",b)
	fmt.Println("a &^ b :::::",a &^ b)
	fmt.Println("a << 1 :::::",a << uint(1))
	fmt.Println("a >> 1 :::::",a >> uint(1))

	fmt.Printf("MULTI-FORMAT ::: %d %[1]b %[1]o %[1]x\n",a)

}
