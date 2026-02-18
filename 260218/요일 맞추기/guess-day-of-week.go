package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func dateTimeDiff(m1, d1, m2, d2 int) int {
	date1 := dateTime(m1, d1)
	date2 := dateTime(m2, d2)
	return date2 - date1
}

func dateTime(m, d int) int {
	days_of_month := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var date int
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

	// Please write your code here.
	diff := dateTimeDiff(m1, d1, m2, d2)
	days := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	var rst string
	if idx := diff%7; idx >= 0 {
		rst = days[idx]
	} else {
		rst = days[7+idx]
	}
	
	fmt.Println(rst)
}