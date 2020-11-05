package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	sc        = bufio.NewScanner(os.Stdin)
	n, m      int
	g         [][]int // 연결상태 저장
	condition []int   // 조건의 개수 저장
)

//	위상정렬...이것도 다시...
//	진입 차수가 0인 노드를 큐에 삽입
//	큐에서 꺼내서 연결된 간선 모두 제거
//	진입 차수가 0이 된 다른 정점들을 큐에 넣음
//	반복 -> 큐에서 꺼낸 결과가 위성정렬의 결과
func main() {
	n, m = nextInt(), nextInt()
	g = make([][]int, n+1)
	condition = make([]int, n+1)
	for i := 1; i <= n; i++ {
		g[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		a, b := nextInt(), nextInt()
		g[a] = append(g[a], b)
		condition[b]++
	}

	q := make([]int, 0)
	for i := 1; i <= n; i++ {
		if condition[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		fmt.Print(q[0], " ")
		cur := q[0]
		q = q[1:]

		for i := 0; i < len(g[cur]); i++ {
			next := g[cur][i]
			condition[next]--
			if condition[next] == 0 {
				q = append(q, next)
			}
		}
	}

}
func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
