package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	br                               = bufio.NewReader(os.Stdin)
	bw                               = bufio.NewWriter(os.Stdout)
	n, m, count                      int
	board                            [][]string
	redX, redY, blueX, blueY, hX, hY int
	dPos                             = []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	check                            [10][10][10][10]bool
)

func main() {
	defer bw.Flush()
	fmt.Fscan(br, &n, &m)
	board = make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, m)
		var str string
		fmt.Fscan(br, &str)
		board[i] = strings.Split(str, "")
		for j := 0; j < m; j++ {
			if board[i][j] == "R" { //	빨간공
				redX, redY = i, j
			}
			if board[i][j] == "B" { //	파란공
				blueX, blueY = i, j
			}
			if board[i][j] == "O" { //	종료
				hX, hY = i, j
			}
		}
	}
	find()
	fmt.Fprint(bw, count)
}
func find() {
	q := make([][4]int, 0) //	빨간공, 파란공 좌표
	disQ := make([]int, 0) //	이동거리
	q = append(q, [4]int{redX, redY, blueX, blueY})
	disQ = append(disQ, 0)
	check[redX][redY][blueX][blueY] = true
	for len(q) > 0 {
		rx := q[0][0]
		ry := q[0][1]
		bx := q[0][2]
		by := q[0][3]
		dis := disQ[0]
		q = q[1:]
		disQ = disQ[1:]

		for i := 0; i < 4; i++ {
			newRX := rx
			newRY := ry
			newBX := bx
			newBY := by
			countR, countB := 0, 0
			newDis := dis + 1

			//	빨간색 이동
			for board[newRX+dPos[i].x][newRY+dPos[i].y] != "#" && board[newRX][newRY] != "O" {
				newRX += dPos[i].x
				newRY += dPos[i].y
				countR++
			}

			//	파란색 이동
			for board[newBX+dPos[i].x][newBY+dPos[i].y] != "#" && board[newBX][newBY] != "O" {
				newBX += dPos[i].x
				newBY += dPos[i].y
				countB++
			}

			//	파란색 들어가면 무시
			if newBX == hX && newBY == hY {
				continue
			}
			//	빨간색 들어가면 움직인 횟수 넣고 종료
			if newRX == hX && newRY == hY {
				count = newDis
				return
			}

			//	두 공의 위치가 겹치면
			//	이동거리에서 더 많이 이동한 쪽의 값을 하나씩 빼준다
			//	임의의 방향으로 이동했을 때 겹친다는건
			//	움직이기 전의 배치와 같기 때문
			if newRX == newBX && newRY == newBY {
				if countR > countB {
					newRX -= dPos[i].x
					newRY -= dPos[i].y
				} else {
					newBX -= dPos[i].x
					newBY -= dPos[i].y
				}
			}

			//	이전에 간 곳이 아니면 BFS 탐색
			if !check[newRX][newRY][newBX][newBY] {
				check[newRX][newRY][newBX][newBY] = true
				q = append(q, [4]int{newRX, newRY, newBX, newBY})
				disQ = append(disQ, newDis)
			}
		}
	}
	//	못찾으면 -1
	count = -1
}

type point struct {
	x, y int
}
