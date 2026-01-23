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
	// B: Black
	// W: White
	// G: Gray
	const (
		B int = iota + 1
		W
		G
	)

	// (1)Tile 구조체 정의
	// - 양방향 Linked list 활용
	// - White, Black 색칠 횟수 카운트 포함
	// - curr: 현재 타일 색깔
	type Tile struct {
		wCnt, bCnt, curr int
		prev, next *Tile
	}

	cursor := &Tile{} // cursor: 현재 가리키고 있는 타일 위치
	for _, cmd := range commands {
		var dst *Tile // 가고자 하는 방향의 다음 칸 타일을 가리킴
		var color int // 어떤 색으로 칠할 것인지
		var cnt *int // 증가시킬 색의 cnt에 직접 접근하기 위함

		// (2)방향에 따라 dst, color, cnt를 세팅
		if cmd.dir == "L" { 
			dst = cursor.prev
			cnt = &cursor.wCnt
			color = W
		} else if cmd.dir == "R" {
			dst = cursor.next
			cnt = &cursor.bCnt
			color = B
		}

		// (3)x 만큼 이동하며 색칠 시작
		for i := range cmd.x {
			// 색칠하기: 방향에 따라 색칠 및 cnt++
			cursor.curr = color
			(*cnt)++

			// 타일 옮기기:
			// - 마지막 반복이라면 건너뛰기
			if i == cmd.x-1 {
				continue
			}
			// - 이미 노드가 있다면 이동하기
			if dst != nil {
				// cursor, dst, cnt는 함께 움직인다
				cursor = dst
				if cmd.dir == "L" {
					dst = cursor.prev
					cnt = &cursor.wCnt
				} else if cmd.dir == "R" {
					dst = cursor.next
					cnt = &cursor.bCnt
				}
			} else { // - 가고자 하는 방향이 nil이라면 새로운 노드 만들기
				tmp := &Tile{}
				if cmd.dir == "L" {
					// 노드 만들고 연결하기
					tmp.next = cursor
					cursor.prev = tmp
					cursor = cursor.prev

					// dst, cnt도 따라간다
					dst = cursor.prev
					cnt = &cursor.wCnt
				} else if cmd.dir == "R" {
					tmp.prev = cursor
					cursor.next = tmp
					cursor = cursor.next

					dst = cursor.next
					cnt = &cursor.bCnt
				}	
			}
		}
	}
	// (4)순회하며 개수 새기
	// - 왼쪽 끝으로 이동
	for cursor.prev != nil { 
		cursor = cursor.prev
	}
	// - 오른쪽 끝으로 이동하며 개수 새기
	var w, b, g int
	for cursor != nil {
		// count가 둘다 2 이상이면 현재 색과 상관없이 회색 카운트 증가
		if cursor.wCnt>=2 && cursor.bCnt>=2 {
			g++
		} else if cursor.curr == W {
			w++
		} else if cursor.curr == B {
			b++
		}
		cursor = cursor.next
	}

	// (5)출력하기
	fmt.Print(w, b, g)
}