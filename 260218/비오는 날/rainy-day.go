package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"slices"
	"cmp"
)

type DayInfo struct{
	y, m, d int
	day string
	weather string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	dayInfo := make([]DayInfo, 100)
	for i := 0; i < n; i++ {
		scanner.Scan()
		fmt.Fscanf(strings.NewReader(scanner.Text()), "%4d-%2d-%2d\n", 
			&dayInfo[i].y, &dayInfo[i].m, &dayInfo[i].d)
		scanner.Scan()
		dayInfo[i].day = scanner.Text()
		scanner.Scan()
		dayInfo[i].weather = scanner.Text()
	}

	// Please write your code here.
	slices.SortStableFunc(dayInfo, func(a, b DayInfo) int {
		return cmp.Compare(a.d, b.d)
	})
	slices.SortStableFunc(dayInfo, func(a, b DayInfo) int {
		return cmp.Compare(a.m, b.m)
	})
	slices.SortStableFunc(dayInfo, func(a, b DayInfo) int {
		return cmp.Compare(a.y, b.y)
	})

	for _, v := range dayInfo {
		if v.weather == "Rain" {
			fmt.Printf("%4d-%02d-%02d %s %s\n", v.y, v.m, v.d, v.day, v.weather)
			break
		}
	}
}