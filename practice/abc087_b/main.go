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
	for iA := 0; iA <= a; iA++ {
		sum += 500 * iA
		for iB := 0; iB <= b; iB++ {
			sum += 100 * iB
			for iC := 0; iC <= c; iC++ {
				sum += 50 * iC
				if sum == x {
					count++
				}
				sum -= 50 * iC
			}
			sum -= 100 * iB
		}
		sum -= 500 * iA
	}
	fmt.Println(count)
}
