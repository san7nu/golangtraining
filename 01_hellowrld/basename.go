package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args[1:]) < 1 {
		fmt.Println("Use Arg")
		os.Exit(1)
	}

	for _, v := range os.Args[1:] {
		st := strings.LastIndex(v, "/")
		end := strings.LastIndex(v, ".")
		if end <= 0 {
			fmt.Println(v[st+1 : ])
		}else {
			fmt.Println(v[st+1 : end])
		}
	}
}
