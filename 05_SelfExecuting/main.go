package main

import "fmt"

func main() {
	func() {
		fmt.Println("This is self executing func")
	}()
}
