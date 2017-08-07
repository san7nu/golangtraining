package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 6}   //Shorthand slice
	b := [10]int{1, 2, 3, 4, 5, 6} //array
	c := make([]int, 6, 10)        //Shorthand slice vth cap defined
	var d []int

	fmt.Println("A:", len(a), cap(a))
	fmt.Println("B:", len(b), cap(b))
	fmt.Println("C:", len(c), cap(c))
	fmt.Println("----------")

	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Println("----------")

	for j := 0; j < len(a); j++ {
		fmt.Print(a[j], ":")
	}
	fmt.Println()
	fmt.Println("----------")

	fmt.Println(a[2:5])
	fmt.Println(a[:5])
	fmt.Println(a[:])
	fmt.Println("SAN"[1])
	fmt.Println("----------")

	//d[1] = 1  This fails
	fmt.Println("D:", d)
	c[1] = 1
	fmt.Println("C:", c[:])
	fmt.Println("----------")

	//Append two slices
	a = append(a, c...) // Observe its not append(a, c)
	fmt.Println("A:", a)
	fmt.Println("----------")

	//To delete use append
	a = append(a[:2], a[5:]...)
	fmt.Println("A:", a)
	fmt.Println("----------")

}
