package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())
	
	// Please write your code here.
	stack := make([]int, 0)
	for {
		if n < b {
			stack = append(stack, n)
			break
		}
		stack = append(stack, n%b)
		n /= b
	}

	for range len(stack) {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(top)
	}
}