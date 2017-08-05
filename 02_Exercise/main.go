package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Print("Enter small num: ")
	fmt.Scan(&num1)
	fmt.Print("Enter big num: ")
	fmt.Scan(&num2)

	fmt.Println("Rem is:", num2%num1)
}
