package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Less(i, j int) bool { return p[i] < p[j] }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {

	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	//Method 1: String sort
	if !sort.StringsAreSorted(studyGroup) {
		//sort.Strings(studyGroup)
	}
	fmt.Println(studyGroup)

	//Method 1: Int sort
	if !sort.IntsAreSorted(n) {
		//sort.Ints(n)
	}
	fmt.Println(n)

	//Method 2: String sort(Interface way)
	//sort.Sort(sort.StringSlice(studyGroup))
	fmt.Println(studyGroup)

	//Method 2: Int sort(Interface way)
	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)

	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}
