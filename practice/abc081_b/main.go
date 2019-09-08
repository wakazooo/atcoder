package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	input := make([]string, 0)
	for i := 0; i < 2; i++ {
		s.Scan()
		input = append(input, s.Text())
	}
	n, _ := strconv.Atoi(input[0])
	arrInput := make([]int, n)
	for idx, i := range strings.Split(input[1], " ") {
		arrInput[idx], _ = strconv.Atoi(i)
	}
	cnt := 0
	for {
		arrInput = converter(arrInput)
		if arrInput != nil {
			cnt++
		} else {
			break
		}
	}
	fmt.Println(cnt)
}

func converter(in []int) []int {
	next := make([]int, len(in))
	for idx, i := range in {
		if i%2 == 0 {
			next[idx] = i / 2
		} else {
			return nil
		}
	}
	return next
}
