package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	// Please write your code here.
	var gcd, lcm int
	for i := 1; i <= n; i++ {
		if n%i==0 && m%i==0 {
			gcd = i
		}
	}
	lcm = n*m/gcd
	fmt.Print(lcm)
}