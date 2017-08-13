package main

import (
	"fmt"
	"sort"
)

func main() {

	type people []string
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Printf("%T \n", studyGroup)
	fmt.Println(sort.Strings([]string(studyGroup)))
}
