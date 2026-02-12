package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func f(n int) {
	if n == 0 {
		return
	}
	for range n {
		fmt.Print("* ")
	}
	fmt.Println()
	f(n-1)
	for range n {
		fmt.Print("* ")
	}
	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nStr := scanner.Text()
	n, _ := strconv.Atoi(nStr)
	// Please write your code here.
	f(n)
}