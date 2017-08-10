package main

import "fmt"

func main() {

	a := make(map[int]string)

	for i := 0; i < 10; i++ {
		a[i] = fmt.Sprintf("%d", i)
	}

	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], ":")
	}
	fmt.Println("\n----------")

	b := make(map[string]string)

	b["a"] = "one"
	b["b"] = "two"

	fmt.Println(b)
	delete(b, "a")
	delete(b, "z")
	fmt.Println(b)
	fmt.Println("----------")

	v, e := b["b"]
	fmt.Println(v,"-",e)

}
