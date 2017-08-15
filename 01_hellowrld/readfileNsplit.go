package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	if len(os.Args[1:]) >= 1 {
		b, _ := ioutil.ReadFile(os.Args[1])
		for _, v := range strings.Split(string(b), "\n") {
			fmt.Println(v)
		}
	}
}
