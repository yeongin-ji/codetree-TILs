package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func f(n, m int) {
	for range n {
		for range m {
			fmt.Print(1)
		}
		fmt.Println()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	// Please write your code here.
	f(n, m)
}