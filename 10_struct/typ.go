package main

import "fmt"

type abc int

type stt struct {
	first string
	last  string
	age   int
}

func main() {

	var a abc
	a = 33
	fmt.Printf("%T %v \n", a, a)
	fmt.Println("----------")

	var p1 stt
	p1.first = "one"
	p1.last = "two"
	p1.age = 30

	p2 := stt{"three", "four", 34}

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("----------")

	fullname1(p1)
	p1.fullname2()
	fmt.Println("----------")

}

func fullname1(p stt) {
	fmt.Println("INSIDE FUNC 11111")
	fmt.Println(p.first, p.last, p.age)
}

// Receiver to a func
func (p stt) fullname2() {
	fmt.Println("INSIDE FUNC 22222")
	fmt.Println(p.first, p.last, p.age)
}
