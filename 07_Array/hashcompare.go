package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("X"))

	fmt.Printf("%b\n%b\n", a, b)
	for i:=0; i<32;i++ {
		fmt.Print(a[i] == b[i]," ")
		fmt.Println(a[i],b[i])
	}

}
