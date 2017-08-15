package main

import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Println("No URL")
		os.Exit(1)
	}
	fmt.Println(os.Args[1])
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println("Some prob vth GET", os.Stderr)
		os.Exit(2)
	}
	
	//First method
	s := bufio.NewScanner(resp.Body)
	defer resp.Body.Close()

	for s.Scan() {
		fmt.Println("== ",s.Text())
	}

	/* //Second method
		b, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println("Some prob vth READALL",os.Stderr)
			os.Exit(2)
		}
	fmt.Println(string(b))
	*/
}
