package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc *bufio.Scanner

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc = bufio.NewScanner(os.Stdin)

	input := nextLine()
	arrInput := strings.Split(input, " ")
	n, _ := strconv.Atoi(arrInput[0])
	a, _ := strconv.Atoi(arrInput[1])
	b, _ := strconv.Atoi(arrInput[2])
	answer := 0
	for i := 1; i <= n; i++ {
		sum := 0
		strN := strconv.Itoa(i)
		arrN := strings.Split(strN, "")
		for j := 0; j < len(arrN); j++ {
			intN, _ := strconv.Atoi(arrN[j])
			sum += intN
		}
		if sum >= a && sum <= b {
			answer += i
		}
	}
	fmt.Println(answer)
}
