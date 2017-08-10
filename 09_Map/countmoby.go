package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	res, _ := http.Get("http://www.gutenberg.org/files/2701/2701-0.txt")
	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	scanner.Split(bufio.ScanWords)

	cntSlice := make([]int, 26)
	var count uint8
	var first string

	for scanner.Scan() {
		first = scanner.Text()
		count = first[0] % 26
		cntSlice[count]++
	}

	for i, v := range cntSlice {
		fmt.Println(string(i+65), " = ", v)
	}
}
