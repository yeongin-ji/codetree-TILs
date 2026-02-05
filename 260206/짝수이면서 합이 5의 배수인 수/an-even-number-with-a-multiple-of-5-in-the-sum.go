package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isEven(n int) bool {
	return n%2==0
}

func isMultiple5(n int) bool {
	a := n/10
	b := n%10
	return (a+b)%5==0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	
	// Please write your code here.
	if isEven(n) && isMultiple5(n) {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}