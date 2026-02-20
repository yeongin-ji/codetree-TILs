package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	tokens := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(tokens[0])
	m, _ := strconv.Atoi(tokens[1])
	k, _ := strconv.Atoi(tokens[2])
	students := make([]int, n+10)
	found := false
	var rst, pStNum int
	for i := 0; i < m; i++ {
		scanner.Scan()
		pStNum, _ = strconv.Atoi(scanner.Text())
		students[pStNum]++
		if !found && students[pStNum] >= k {
			rst = pStNum
			found = true
		}
	}
	if found {
		fmt.Print(rst)
	} else {
		fmt.Print(-1)
	}
	// Please write your code here.
}