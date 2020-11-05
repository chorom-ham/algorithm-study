package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	sc          = bufio.NewScanner(os.Stdin)
	tc, n, m, w int
)

//	이해가 안됨
//	다시 풀어볼것
func main() {
	tc = nextInt()
	for i := 0; i < tc; i++ {
		n, m, w = nextInt(), nextInt(), nextInt()
		edges := make([]edge, 0)
		for j := 0; j < m; j++ { //	도로 입력, 방향 x
			s, e, t := nextInt(), nextInt(), nextInt()
			edges = append(edges, edge{s, e, t})
			edges = append(edges, edge{e, s, t})
		}
		for j := 0; j < w; j++ { //	웜홀 입력, 방향 o
			s, e, t := nextInt(), nextInt(), nextInt()
			edges = append(edges, edge{s, e, (-1) * t})
		}

		dis := make([]int, n+1) //	시작지점에서의 최단 경로
		for j := range dis {
			dis[j] = math.MaxInt32
		}
		dis[1] = 0        //	시작지점
		isUpdate := false //	업데이트가 일어나는지에 대한 판단

		//	edge relocation
	loop:
		for j := 1; j <= n; j++ {
			isUpdate = false //	매번 relocation 할때마다 false
			for _, edge := range edges {
				if dis[edge.end] > dis[edge.start]+edge.cost {
					dis[edge.end] = dis[edge.start] + edge.cost
					isUpdate = true //	업데이트 되면 true

					//	n-1번째 까지 해야하지만 이후에도 계속 루프가 돈다면
					//	업데이트가 계속 되고 있다는 것
					//	즉 negative cycle을 가지고 있다
					if i == n {
						isUpdate = true
						break loop
					}
				}
			}
			if !isUpdate { //	relocation이 끝났는데 false면 업데이트가 일어나지 않았으므로 종료
				break
			}
		}
		if isUpdate {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

type edge struct {
	start, end, cost int
}

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
