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

	firstInput := nextLine()
	a, _ := strconv.Atoi(firstInput)

	secondInput := strings.Split(nextLine(), " ")
	b, _ := strconv.Atoi(secondInput[0])
	c, _ := strconv.Atoi(secondInput[1])
	s := nextLine()

	fmt.Println(fmt.Sprintf("%d", a+b+c), s)
}
