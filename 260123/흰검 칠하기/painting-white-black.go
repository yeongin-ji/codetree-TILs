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

	readString := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	readInt := func() int {
		num, _ := strconv.Atoi(readString())
		return num
	}

	N := readInt()
	commands := make([]struct {
		x   int
		dir string
	}, N)
	for i := 0; i < N; i++ {
		commands[i].x = readInt()
		commands[i].dir = readString()
	}

	// Please write your code here.
	const (
		B int = iota + 1
		W
		G
	)

	type Tile struct {
		wCnt, bCnt int
		curr int
		next *Tile
		prev *Tile
	}

	t := &Tile{}
	cursor := t
	for _, cmd := range commands {
		//fmt.Println(j, "번째 명령을 실행합니다.")
		// 방향, 색 설정
		var dst *Tile
		var color int
		var cnt *int
		if cmd.dir == "L" { 
			dst = cursor.prev
			cnt = &cursor.wCnt
			color = W
		} else if cmd.dir == "R" {
			dst = cursor.next
			cnt = &cursor.bCnt
			color = B
		}
		for i := range cmd.x {
			// 색칠하기: 회색이 아니라면 방향에 따라 색칠
			if cursor.curr != G {
				cursor.curr = color
				*cnt = *cnt + 1
			}

			// 타일 옮기기 (왼/오)
			// (1) 마지막 반복이라면 건너뛰기
			if i == cmd.x-1 {
				continue
			}
			// (2) 이미 노드가 있다면 이동하기
			// (3) 가고자 하는 방향이 nil이라면 새로운 노드 만들기
			if dst != nil {
				//fmt.Println("진행 방향에 타일이 이미 존재합니다.")
				cursor = dst
				if cmd.dir == "L" {
					dst = cursor.prev
					cnt = &cursor.wCnt
					//fmt.Println("왼쪽으로 이동합니다.")
				} else if cmd.dir == "R" {
					dst = cursor.next
					cnt = &cursor.bCnt
					//fmt.Println("오른쪽으로 이동합니다.")
				}
			} else {
				//fmt.Println("진행 방향에 타일이 없습니다. 새로 만듭니다.")
				tmp := &Tile{}
				if cmd.dir == "L" {
					tmp.next = cursor
					cursor.prev = tmp
					cursor = cursor.prev
					//fmt.Println("왼쪽으로 이동합니다.")

					dst = cursor.prev
					cnt = &cursor.wCnt
				} else if cmd.dir == "R" {
					tmp.prev = cursor
					cursor.next = tmp
					cursor = cursor.next
					//fmt.Println("오른쪽으로 이동합니다.")

					dst = cursor.next
					cnt = &cursor.bCnt
				}	
			}
			//fmt.Println()
		}
	}
	// 순회하며 개수 새기
	// (1) 왼쪽 끝으로 이동
	for cursor.prev != nil { 
		cursor = cursor.prev
	}
	// (2) 오른쪽 끝으로 이동하며 개수 새기
	var w, b, g int
	for cursor != nil {
		if cursor.wCnt>=2 && cursor.bCnt>=2 {
			g++
		} else if cursor.curr == W {
			w++
		} else if cursor.curr == B {
			b++
		}
		cursor = cursor.next
	}

	// 출력하기
	fmt.Print(w, b, g)
}