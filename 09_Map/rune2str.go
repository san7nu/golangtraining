package main

import "fmt"

func main() {

	fmt.Println(" ------ RUNE ----- ")
	v1 := 'A'
	fmt.Println(v1)
	fmt.Printf("%T \n", v1)

	fmt.Println(" ------ STRING TO RUNE ----- ")
	v2 := "A"
	fmt.Println(v2[0])
	fmt.Printf("%T \n", v2[0])
}
