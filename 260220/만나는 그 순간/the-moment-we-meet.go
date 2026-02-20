package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type move struct {
	d string
	t int
}

func Simul(mvHst [][2]int, mv []move, id int) int {
	currT := 0
	currPos := 0
	for _, m := range mv {
		var sign int
		if m.d == "L" {
			sign = -1
		} else if m.d == "R" {
			sign = 1
		}
		for range m.t {
			currT++
			currPos += sign
			mvHst[currT][id] = currPos
		}
	}
	return currT
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	movesA := make([]move, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		d := scanner.Text()
		scanner.Scan()
		t, _ := strconv.Atoi(scanner.Text())
		movesA[i] = move{d, t}
	}

	movesB := make([]move, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		d := scanner.Text()
		scanner.Scan()
		t, _ := strconv.Atoi(scanner.Text())
		movesB[i] = move{d, t}
	}

	// Please write your code here.
	moveHist := make([][2]int, 1100)
	var totalT int
	totalT = Simul(moveHist, movesA, 0)
	totalT = Simul(moveHist, movesB, 1)
	
	met := false
	for i := range totalT {
		if moveHist[i+1][0]==moveHist[i+1][1] {
			fmt.Print(i+1)
			met = true
			break
		}
	}
	if !met {
		fmt.Print(-1)
	}
}