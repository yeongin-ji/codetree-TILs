package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func f(a, b int) {
	prod := 1
	for range b {
		prod *= a
	}
	fmt.Print(prod)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	
	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())
	
	// Please write your code here.
	f(a, b)
}