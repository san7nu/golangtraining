package main

import "fmt"

func main() {
	a := []string{"abc", "def", "def", "xyz", "xyz"}
	fmt.Println(a)
	remD(a)
	fmt.Println(a)
}

func remD(n []string) {
	for i := 1; i < len(n); i++ {
		if len(n[i-1]) == len(n[i]) {
			j := 0
			for j = 0; j < len(n[i-1]); j++ {
				if n[i-1][j] != n[i][j] {
					break
				}
			}
			if j == len(n[i]) {
				copy(n[i-1:], n[i:])
				n[len(n)-1] = ""
			}
		}
	}
}
