package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Args[1:])

	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
	fmt.Println("==============")

	for i, v := range os.Args {
		fmt.Println(i, "=", v)
	}
	fmt.Println("==============")
}
