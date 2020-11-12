package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc     = bufio.NewScanner(os.Stdin)
	n, num int
	min    = math.MaxInt32
	maap   = [100][100]int{}
	dis    = [100][100]int{}
	dx     = []int{-1, 1, 0, 0}
	dy     = []int{0, 0, -1, 1}
	check  = [100][100]bool{}
	isEnd  bool
)

func main() {
	n = nextInt()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			maap[i][j] = nextInt()
			dis[i][j] = math.MaxInt32
		}
	}

	//	섬 번호 붙이기
	num = 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !check[i][j] && maap[i][j] == 1 {
				setNumber(i, j)
				num++
			}
		}
	}

	//	다리계산
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if maap[i][j] != 0 {
				find(maap[i][j], i, j, 0)
			}
		}
	}
	fmt.Print(min)
}

//	다리 연결, BFS 이용
func find(landNum, x, y, count int) {
	if min <= count {
		return
	}
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]

		if checkRange(nx, ny) {
			continue
		}
		if dis[nx][ny] <= count+1 {
			continue
		}
		if maap[nx][ny] == 0 {
			dis[nx][ny] = count + 1
			find(landNum, nx, ny, count+1)
			continue
		}
		if maap[nx][ny] != landNum {
			min = minInt(count, min)
			return
		}
	}
}

//	번호 붙이기, DFS 이용
func setNumber(x, y int) {
	check[x][y] = true
	maap[x][y] = num
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if !checkRange(nx, ny) && !check[nx][ny] && maap[nx][ny] != 0 {
			setNumber(nx, ny)
		}
	}
}

func checkRange(x, y int) bool {
	return (x < 0 || x >= n || y < 0 || y >= n)
}
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type pos struct {
	x, y int
}

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
