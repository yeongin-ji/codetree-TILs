package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	for i := 2; i <= n-1; i++ {
		if n%i==0 {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	// Please write your code here.
	sum := 0
	for i := a; i <= b; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Print(sum)
}