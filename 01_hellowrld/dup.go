package main

import (
	"fmt"
	"strings"
)

func main() {
	var en []string
	var a string

	fmt.Print("Enter in(Z is end char): ")
	for a != string('Z') {
		fmt.Scan(&a)
		en = append(en, a)
	}
	//fmt.Println("===>", en)
	fmt.Println("Duplicate lines are:")

	//Logic to find duplicate lines: Find equal length lines then compare if thy r equal

	for i := 0; i < len(en); i++ {
		for j := i+1; j < len(en); j++ {
			if len(en[i]) == len(en[j]) {
				if strings.EqualFold(en[i], en[j]) {
					fmt.Println(en[i])
				}
			}
		}
	}
}
