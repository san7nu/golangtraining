package main

import (
	"encoding/json"
	"fmt"
)


func main() {

	type str struct {
		Id   int
		Name string
	}

	a := str{1, "abc"}
	fmt.Println(a)
	
	b, _ := json.Marshal(a)
	fmt.Println(string(b))

	var c str

	_ = json.Unmarshal(b, &c)
	fmt.Println(c)
}
