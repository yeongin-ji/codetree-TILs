package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ConvAto10(a int, n string) int {
	var rst int
	for _, r := range n {
		rst = rst*a + int(r-'0')
	}
	return rst
}

// n을 b진수로 바꿈
func Conv10toB(b, n int) string {
	st := make([]int, 0)
	for {
		if n < b {
			st = append(st, n)
			break
		}
		st = append(st, n%b)
		n /= b
	}
	var sb strings.Builder
	for range len(st) {
		top := st[len(st)-1]
		c := fmt.Sprintf("%d", top)
		sb.WriteString(c)
		st = st[:len(st)-1]
	}
	return sb.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	aStr, bStr := "", ""
	fmt.Sscan(scanner.Text(), &aStr, &bStr)
	a, _ := strconv.Atoi(aStr)
	b, _ := strconv.Atoi(bStr)

	scanner.Scan()
	n := scanner.Text()

	// Please write your code here.
	conv := ConvAto10(a, n)
	rst := Conv10toB(b, conv)
	fmt.Print(rst)
}