package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isMagicNum(n int) bool {
	return isPrime(n) && isEven(n)
}

func isPrime(n int) bool {
	for i := 2; i <= n-1; i++ {
		if n%i==0 {
			return false
		}
	}
	return true
}

func isEven(n int) bool {
	s := strconv.Itoa(n)
	sum := 0
	for _, r := range s {
		sum += int(r-'0')
	}
	if sum%2==0 {
		return true
	} else {
		return false
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	// Please write your code here.
	cnt := 0
	for i := a; i <= b; i++ {
		if isMagicNum(i) {
			cnt++
		}
	}
	fmt.Print(cnt)
}