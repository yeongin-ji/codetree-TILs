package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func f(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())

	// Please write your code here.
	min := f(a, b, c)
	fmt.Print(min)
}