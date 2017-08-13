package main

import (
	"encoding/json"
	"fmt"
)

type stt struct {
	First string
	Last  string 
	//Age   int   `json:"comment"`
	Age   int   
}

func main() {
	a := stt{"one", "two", 12}
	bs, _ := json.Marshal(a)

	fmt.Println(a)
	fmt.Println(string(bs))
	fmt.Println("------------")

	var b stt
	byt := []byte(`{"First":"James", "Last":"Bond", "Age":20}`)

	json.Unmarshal(byt,&b)
	fmt.Println(b)
	fmt.Println("------------")
}
