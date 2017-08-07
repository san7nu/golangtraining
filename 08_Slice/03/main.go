package main

import "fmt"

func main() {

	a := []int{}
	var b []int
	c := make([]int, 0)   // Lets use this style 4m now on

	// Does declaration make thm nill ?
	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(c == nil)
	fmt.Println("---------")

	// Can i append any type of declartn ;         Seems i cannot assign(a[5] = 5) beyond wht is declared
	for i := 0; i < 10; i++ {
		a = append(a, i)
		b = append(b, i)
		c = append(c, i)
	}
	fmt.Println("A:", a)
	fmt.Println("B:", b)
	fmt.Println("C:", c)
	fmt.Println("---------")

	//N-dimension
	two := make([][]int,0)
	part1 := make([]int, 2)
	part2 := make([]int, 2)

	part1[0] = 11
	part1[1] = 22

	part2[0] = 13
	part2[1] = 26

	two = append(two, part1)     // Its not append(two, part1...)
	two = append(two, part2)
	fmt.Println("TWO:", two)
	fmt.Println("---------")

}
