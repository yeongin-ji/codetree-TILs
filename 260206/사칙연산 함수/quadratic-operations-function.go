package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(a, b int) {
	fmt.Printf("%d + %d = %d", a, b, a+b)
}

func sub(a, b int) {
	fmt.Printf("%d - %d = %d", a, b, a-b)
}

func div(a, b int) {
	fmt.Printf("%d / %d = %d", a, b, a/b)
}

func mul(a, b int) {
	fmt.Printf("%d * %d = %d", a, b, a*b)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	aStr := scanner.Text()
	scanner.Scan()
	op := scanner.Text()
	scanner.Scan()
	cStr := scanner.Text()

	a, _ := strconv.Atoi(aStr)
	c, _ := strconv.Atoi(cStr)

	// Please write your code here.
	switch op {
	case "+":
		sum(a, c)
	case "-":
		sub(a, c)
	case "/":
		div(a, c)
	case "*":
		mul(a, c)
	default:
		fmt.Print("False")
	}
}