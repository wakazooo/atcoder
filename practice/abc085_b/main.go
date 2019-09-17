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
	n, _ := strconv.Atoi(nextLine())
	args := args{
		arr: make([]int, 0, n),
	}
	for index := 0; index < n; index++ {
		input, _ := strconv.Atoi(nextLine())
		sort(&args, input)
	}
	ans := calc(args)
	fmt.Println(ans)
}

type args struct {
	arr []int
}

func sort(args *args, input int) *args {
	if len(args.arr) == 0 {
		args.arr = append(args.arr, input)
		return args
	}
	index := 0
	for ; index < len(args.arr); index++ {
		head := args.arr[index]
		if head == input {
			return args
		} else if head < input {
			break
		}
	}
	if index == len(args.arr) {
		args.arr = append(args.arr, input)
	} else {
		args.arr = append(args.arr[:index+1], args.arr[index:]...)
		args.arr[index] = input
	}
	return args
}

func calc(args args) int {
	return len(args.arr)
}
