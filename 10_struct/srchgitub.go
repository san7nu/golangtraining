package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	//	"io/ioutil"
)

func main() {

	type usr struct {
		Login string `json:"login"`
	}
	type itm struct {
		Url   string `json:"url"`
		Users *usr
	}
	type respStru struct {
		Count int `json:"total_count"`
		Items []itm
	}

	resp, err := http.Get("https://api.github.com/search/issues?q=abc")
	fmt.Println(err)
	defer resp.Body.Close()
	/*
	   	page, _ := ioutil.ReadAll(resp.Body)
	           fmt.Println(string(page))
	*/

	var c respStru

	err = json.NewDecoder(resp.Body).Decode(&c)
	fmt.Println(err)
	
	fmt.Printf("%T\n",c)
	fmt.Println(c)

}
