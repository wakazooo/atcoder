package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	r := 0
	for _, i := range input {
		if int(i-'0') == 1 {
			r += int(i - '0')
		}
	}
	fmt.Println(r)
}
