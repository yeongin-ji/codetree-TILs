package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func f(n int) {
	for range n {
		fmt.Printf("12345^&*()_\n")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	rowNum, _ := strconv.Atoi(scanner.Text())
	// Please write your code here.
	f(rowNum)
}