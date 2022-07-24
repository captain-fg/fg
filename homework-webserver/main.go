package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%s is starting to run his code.", "Curry Fu")
	var sli = make([]string, 10, 10)
	for i := 0; i < 10; i++ {
		sli[i] = strconv.Itoa(i)
	}
	fmt.Println(sli)

}
