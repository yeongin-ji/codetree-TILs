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
	a, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	d, _ := strconv.Atoi(scanner.Text())

	// Please write your code here.
	time1 := a*60 + b
	time2 := c*60 + d
	rst := time2 - time1
	fmt.Println(rst)
}