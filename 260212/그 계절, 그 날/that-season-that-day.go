package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isYoon(y int) bool {
	if y%4==0 {
		if y%400==0 {
			return true
		}
		if y%100==0 {
			return false
		}
		return true
	} else {
		return false
	}
}

func isExist(y, m, d int) bool {
	switch m {
	case 4, 6, 9, 11:
		if d > 30 {
			return false
		}
	case 2:
		if d > 28 {
			if d==29 && isYoon(y) {
				return true
			} else {
				return false
			}
		} 
	default:
		if d > 31 {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	yStr := scanner.Text()
	scanner.Scan()
	mStr := scanner.Text()
	scanner.Scan()
	dStr := scanner.Text()

	y, _ := strconv.Atoi(yStr)
	m, _ := strconv.Atoi(mStr)
	d, _ := strconv.Atoi(dStr)

	// Please write your code here.
	if isExist(y, m, d) {
		switch m {
		case 3, 4, 5:
			fmt.Print("Spring")
		case 6, 7, 8:
			fmt.Print("Summer")
		case 9, 10, 11:
			fmt.Print("Fall")
		case 12, 1, 2:
			fmt.Print("Winter")
		}
	} else {
		fmt.Print(-1)
	}
}