package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

	input1 := nextLine()
	N, _ := strconv.Atoi(input1)
	input2 := nextLine()
	arrStr := strings.Split(input2, " ")
	arrInt := make([]int, N)
	for i, a := range arrStr {
		arrInt[i], _ = strconv.Atoi(a)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arrInt)))
	alice := 0
	bob := 0
	for index, a := range arrInt {
		if index%2 == 0 {
			alice += a
		} else {
			bob += a
		}
	}
	ans := alice - bob
	fmt.Println(math.Abs(float64(ans)))
}
