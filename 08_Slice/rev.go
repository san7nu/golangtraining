package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 8, 7, 9}
	fmt.Println(a)
	rev(a)
	fmt.Println(a)


}

func rev(n []int) {
	for i:=0;i<(len(n)/2);i++ {
		n[i],n[len(n)-i-1] = n[len(n)-i-1],n[i]
	}
}
