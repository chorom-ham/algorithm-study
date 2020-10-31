package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var (
	br          = bufio.NewReader(os.Stdin)
	bw          = bufio.NewWriter(os.Stdout)
	w, h, count int
	maap        [][]string
	laser       []point
	check       [100][100]int
	dPos        = []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
)

//	다시 풀것
func main() {
	defer bw.Flush()
	fmt.Fscan(br, &w, &h)
	maap = make([][]string, h)
	for i := 0; i < h; i++ {
		var str string
		fmt.Fscan(br, &str)
		maap[i] = make([]string, w)
		maap[i] = strings.Split(str, "")
		for j := 0; j < w; j++ {
			if maap[i][j] == "C" { //	레이저 시작, 끝
				laser = append(laser, point{i, j})
			}
			check[i][j] = math.MaxInt32
		}
	}
	find()
	fmt.Fprint(bw, count)
}
func find() {
	startX, startY, endX, endY := laser[0].x, laser[0].y, laser[1].x, laser[1].y
	q := make([][4]int, 4)
	for i := range q {
		q = append(q, [4]int{startX, startY, i, 0})
	}
	//	시작점은 거울 필요x
	check[startX][startY] = 0
	for len(q) > 0 {
		x := q[0][0]
		y := q[0][1]
		pos := q[0][2]
		mir := q[0][3]
		q = q[1:]

		for i := 0; i < 4; i++ {
			newX := x + dPos[i].x
			newY := y + dPos[i].y
			newM := mir

			//	범위 넘어가면 패스
			if newX < 0 || newY < 0 || newX >= h || newY >= w {
				continue
			}
			//	벽이 아닐 때
			if maap[newX][newY] != "*" {
				//	현재 진행방향과 같지 않으면 거울 갯수 추가
				if pos != i {
					newM++
				}
				//	방문하지 않았다면, 거울 갯수로 체크
				if check[newX][newY] >= newM {
					check[newX][newY] = newM
					q = append(q, [4]int{newX, newY, i, newM})
				}
			}
		}
	}
	//	탐색이 끝나면 종료지점 거울 갯수 대입
	count = check[endX][endY]
}

type point struct {
	x, y int
}
