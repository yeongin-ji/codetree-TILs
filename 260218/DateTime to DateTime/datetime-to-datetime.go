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
	
	// Please write your code here.
	dateTime := (60*24*a+60*b+c) - (60*24*11+60*11+11)
	if dateTime >= 0 {
		fmt.Println(dateTime)
	} else {
		fmt.Println(dateTime)
	}
}