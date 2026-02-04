package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func f(n, m int) {
	var t []int = make([]int, 0, 100)
	for i := 1; i <= n; i++ {
		if n%i==0 {
			t = append(t, i)
		}
	}
	gbm := 1
	for _, v := range t {
		if m%v==0 {
			gbm = v
		}
	}
	fmt.Print(gbm)
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