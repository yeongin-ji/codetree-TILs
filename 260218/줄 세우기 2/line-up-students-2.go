package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
	"cmp"
)

type Student struct {
	height int
	weight int
	number int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	students := make([]Student, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		height, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		weight, _ := strconv.Atoi(scanner.Text())
		students[i] = Student{height, weight, i + 1}
	}

	// Please write your code here.
	slices.SortFunc(students, func(a, b Student) int {
		if n := cmp.Compare(a.height, b.height); n!=0 {
			return n
		}
		return cmp.Compare(b.weight, a.weight)
	})

	for _, st := range students {
		fmt.Printf("%d %d %d\n", st.height, st.weight, st.number)
	}
}