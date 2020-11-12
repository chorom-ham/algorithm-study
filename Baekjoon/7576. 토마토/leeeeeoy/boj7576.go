package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var (
	tomato      [1000][1000]int
	day         [1000][1000]int
	n, m, count int
	dx          = [4]int{1, 0, -1, 0}
	dy          = [4]int{0, 1, 0, -1}
	q           = make([]pos, 0)
)

func main() {
	m, n = nextInt(), nextInt()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tomato[i][j] = nextInt()
			if tomato[i][j] == 1 {
				q = append(q, pos{i, j})
			}
		}
	}
	fmt.Println(q)
	find()
	isAll()
}
func find() {
	for len(q) > 0 {
		xx := q[0].x
		yy := q[0].y
		q = q[1:]
		for i := 0; i < 4; i++ {
			xxx := xx + dx[i]
			yyy := yy + dy[i]
			if xxx >= 0 && yyy >= 0 && xxx < n && yyy < m {
				if day[xxx][yyy] == 0 && tomato[xxx][yyy] == 0 {
					q = append(q, pos{xxx, yyy})
					day[xxx][yyy] = day[xx][yy] + 1
				}
			}
		}
	}
}
func isAll() {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			count = maxInt(count, day[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if day[i][j] == 0 && tomato[i][j] == 0 {
				fmt.Print(-1)
				return
			}
		}
	}
	fmt.Print(count)
}

type pos struct {
	x int
	y int
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
