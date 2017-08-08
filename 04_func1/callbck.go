package main

import "fmt"

func abc(n int, zzz func()) {
	fmt.Println("---->", n)
	zzz()
}

func main() {

	abc(10, func() {
		fmt.Println("INSIDE")
	})

}
