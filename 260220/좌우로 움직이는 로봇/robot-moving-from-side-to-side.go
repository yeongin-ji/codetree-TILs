package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type move struct {
	t int
	d string
}

const size = 1_000_000

func SetMoves(hist []int, moves []move) int {
	pos := 0
	totalTime := 0
	for _, m := range moves {
		pos, time = SetMoveOneCmd(hist, m, pos, time)
		totalTime += m.t
	}
	return totalTime
}

func SetMoveOneCmd(h []int, m move, pos, time int) int, int {
	var dir int
	if m.d == "L" {
		dir = -1
	} else if m.d == "R" {
		dir = 1
	}
	for i := 1; i <= m.t; i++ {
		pos += dir
		h[time+i] = pos
	}
	return pos, time+m.t
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	
	n, _ := strconv.Atoi(nextToken(scanner))
	m, _ := strconv.Atoi(nextToken(scanner))
	
	robotAMoves := make([]move, n)
	for i := 0; i < n; i++ {
		t, _ := strconv.Atoi(nextToken(scanner))
		d := nextToken(scanner)
		robotAMoves[i] = move{t, d}
	}
	
	robotBMoves := make([]move, m)
	for i := 0; i < m; i++ {
		t, _ := strconv.Atoi(nextToken(scanner))
		d := nextToken(scanner)
		robotBMoves[i] = move{t, d}
	}
	
	// Please write your code here.
	aHist := make([]int, size+100)
	bHist := make([]int, size+100)
	ATotalTime := SetMoves(aHist, robotAMoves)
	BTotalTime := SetMoves(bHist, robotBMoves)
	var longTime int
	var bLong, aLong bool
	if ATotalTime <= BTotalTime {
		longTime = BTotalTime
		// 마지막 원소로 채워서 길이 맞추기
		for i := ATotalTime+1; i <= BTotalTime; i++ {
			aHist[i] = aHist[ATotalTime]
		}
	} else {
		longTime = ATotalTime
		for i := BTotalTime+1; i <= ATotalTime; i++ {
			bHist[i] = bHist[BTotalTime]
		}
	}
	cnt := 0
	for i := 2; i <= longTime; i++ {
		currEncounted := aHist[i] == bHist[i]
		prevEncounted := aHist[i-1] == bHist[i-1]
		if !prevEncounted && currEncounted {
			cnt++
		}
	}
	fmt.Print(cnt)
}

func nextToken(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}