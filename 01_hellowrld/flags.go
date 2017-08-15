package main

import (
	"flag"
	"fmt"
)

func main() {
	i := flag.String("n", "test", "testing flag")

	fmt.Println(*i)
}
