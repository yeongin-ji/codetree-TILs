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
	m1, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	d1, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m2, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	d2, _ := strconv.Atoi(scanner.Text())

	// Please write your code here.
	days_of_month := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var date1, date2 int
	for i := range m1 {
		date1 += days_of_month[i+1]
	}
	date1 += d1
	for i := range m2 {
		date2 += days_of_month[i+1]
	}
	date2 += d2
	rst := date2 - date1 + 1
	fmt.Println(rst)
}