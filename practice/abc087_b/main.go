package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc *bufio.Scanner

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc = bufio.NewScanner(os.Stdin)

	input1 := nextLine()
	a, _ := strconv.Atoi(input1)

	input2 := nextLine()
	b, _ := strconv.Atoi(input2)

	input3 := nextLine()
	c, _ := strconv.Atoi(input3)

	input4 := nextLine()
	x, _ := strconv.Atoi(input4)
	sum := 0
	count := 0
	for i_a := 0; i_a <= a; i_a++ {
		sum += 500 * i_a
		for i_b := 0; i_b <= b; i_b++ {
			sum += 100 * i_b
			for i_c := 0; i_c <= c; i_c++ {
				sum += 50 * i_c
				if sum == x {
					count++
				}
				sum -= 50 * i_c
			}
			sum -= 100 * i_b
		}
		sum -= 500 * i_a
	}
	fmt.Println(count)
}
