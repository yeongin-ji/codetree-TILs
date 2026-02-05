package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func f(n int) int {
	rst := 0
	for i := 1; i <= n; i++ {
		rst += i
	}
	rst /= 10
	return rst
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	// Please write your code here.
	_ = n

	rst := f(n)
	fmt.Println(rst)
}