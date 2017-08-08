package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	res, _ := http.Get("https://www.google.com/")
	defer res.Body.Close()
	fmt.Println(res.Status)

	page, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%T \n", page)

}
