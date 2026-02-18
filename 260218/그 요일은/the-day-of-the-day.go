package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func DateTimeDiff(m1, d1, m2, d2 int) int {
	date1 := DateTime(m1, d1)
	date2 := DateTime(m2, d2)
	return date2 - date1
}

func DateTime(m, d int) int {
	var date int
	days_of_month := [13]int{0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for i := 1; i < m; i++ {
		date += days_of_month[i]
	}
	date += d
	return date
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	
	scanner.Scan()
	m1, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	d1, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m2, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	d2, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	A := scanner.Text()
	
	// Please write your code here.
	diff := DateTimeDiff(m1, d1, m2, d2)
	days := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	div := diff/7
	mod := diff%7

	var A_idx int
	for i, d := range days {
		if A == d {
			A_idx = i
		}
	}

	var rst int
	if mod >= A_idx {
		rst = div + 1
	} else {
		rst = div
	}

	fmt.Print(rst)
}