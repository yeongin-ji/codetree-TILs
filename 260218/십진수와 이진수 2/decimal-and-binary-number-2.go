package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	binaryStr := scanner.Text()
	
	// Please write your code here.
	var decimal int
	for _, r := range binaryStr {
		decimal = decimal*2 + int(r-'0')
	}
	decimal *= 17
	stack := make([]int, 0)
	for {
		if decimal < 2 {
			stack = append(stack, decimal)
			break
		}
		stack = append(stack, decimal%2)
		decimal /= 2
	}
	for range len(stack) {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(top)
	}
}