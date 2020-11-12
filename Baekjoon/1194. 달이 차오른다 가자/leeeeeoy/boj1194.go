package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br    = bufio.NewReader(os.Stdin)
	bw    = bufio.NewWriter(os.Stdout)
	n, m  int
	maze  [][]byte
	dx    = []int{-1, 1, 0, 0}
	dy    = []int{0, 0, -1, 1}
	check [50][50][1 << 6]bool
)

func main() {
	defer bw.Flush()
	fmt.Fscan(br, &n, &m)
	maze = make([][]byte, n)
	start, end := 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(br, &maze[i])
		for j := 0; j < m; j++ {
			if maze[i][j] == '0' {
				start, end = i, j
			}
		}
	}
	answer := find(start, end)
	fmt.Fprint(bw, answer)
}
func find(s, e int) int {
	q := make([]node, 0)
	q = append(q, node{s, e, 0, 0})
	check[s][e][0] = true

	for len(q) > 0 {
		cur := q[0]
		curX := cur.x
		curY := cur.y
		curK := cur.key
		curC := cur.count
		q = q[1:]

		if maze[curX][curY] == '1' {
			return curC
		}

		for i := 0; i < 4; i++ {
			nx := curX + dx[i]
			ny := curY + dy[i]

			if nx >= 0 && nx < n && ny >= 0 && ny < m { //	범위 체크
				if maze[nx][ny] != '#' && !check[nx][ny][curK] { //	벽이 아니고 방문 안했을 때
					if maze[nx][ny] == '.' || maze[nx][ny] == '0' || maze[nx][ny] == '1' { //	움직일 수 있는 곳들
						check[nx][ny][curK] = true
						q = append(q, node{nx, ny, curK, curC + 1})
					} else if maze[nx][ny] >= 'a' && maze[nx][ny] <= 'z' { //	열쇠
						nKey := (1 << (maze[nx][ny] - 'a')) | curK //	현재 가지고 있는 열쇠와 다음 열쇠 정보 갱신
						if !check[nx][ny][nKey] {
							check[nx][ny][curK] = true
							check[nx][ny][nKey] = true
							q = append(q, node{nx, ny, nKey, curC + 1})
						}
					} else if maze[nx][ny] >= 'A' && maze[nx][ny] <= 'Z' { //	문
						door := 1 << (maze[nx][ny] - 'A')
						if (door & curK) > 0 { //	문과 열쇠의 AND 연산이 0보다 크면 문을 열 수 있음
							check[nx][ny][curK] = true
							q = append(q, node{nx, ny, curK, curC + 1})
						}
					}
				}
			}
		}
	}
	return -1
}

type node struct {
	x, y, key, count int
}
