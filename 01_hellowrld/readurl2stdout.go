package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	var url string

	if len(os.Args[1:]) < 1 {
		fmt.Println("NO URL")
		os.Exit(1)
	}

	if strings.HasPrefix(os.Args[1], "http://") {
		url = os.Args[1]
	}else {
		url = "http://" + os.Args[1]
	}
	fmt.Println("URL", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Prob vth GET", os.Stderr)
		os.Exit(2)
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Println("Prob vth COPY", os.Stderr)
		os.Exit(3)
	}
	fmt.Println("STATUS CODE", resp.StatusCode)

}
