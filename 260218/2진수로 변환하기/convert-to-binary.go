package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nStr := scanner.Text()
	n, _ := strconv.Atoi(nStr)
	
	// Please write your code here.
	stack := make([]int, 0)
	for  {
		if n == 0 || n == 1 {
			stack = append(stack, n)
			break
		}
		stack = append(stack, n%2)
		n /= 2
	}

	for range len(stack) {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(pop)
	}
}