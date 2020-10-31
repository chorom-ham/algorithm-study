package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var (
	sc                        = bufio.NewScanner(os.Stdin)
	n, m, num, isLand, result int
	parents                   []int
	maap                      [10][10]int
	isEnd                     bool
	check                     [10][10]bool
	dx                        = []int{-1, 1, 0, 0}
	dy                        = []int{0, 0, -1, 1}
	edges                     []edge
)

//	다리 길이 2 이상
//	방향 바뀌면 안됨
func main() {
	n, m = nextInt(), nextInt()
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			maap[i][j] = nextInt()
		}
	}

	//	섬에 넘버링
	num = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !check[i][j] && maap[i][j] == 1 {
				setNumber(i, j, num) //	BFS로 탐색
				num++
			}
		}
	}

	isLand = num

	//	다리 만들기
	edges = make([]edge, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if maap[i][j] >= 1 { //	다리 시작은 섬(땅)부터 시작
				num = maap[i][j]
				for d := 0; d < 4; d++ {
					isEnd = false
					if checkRange(i+dx[d], j+dy[d]) { //	범위 넘어가는지 체크
						continue
					}
					if maap[i+dx[d]][j+dy[d]] == 0 { //	0이면 다리 생성 가능
						makeBridge(i+dx[d], j+dy[d], d, 1) //	DFS로 연결
					}
				}
			}
		}
	}
	//	최소스패닝트리
	//	가중치를 기준으로 정렬
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	parents = make([]int, isLand+1)
	for i := 0; i < len(parents); i++ {
		parents[i] = -1
	}

	cnt := 0
	for _, edge := range edges { //	연결 정보를 하나씩 처리한다
		if union(edge.node1, edge.node2) { //	임의의 정점이 연결되어 있는지 확인
			result += edge.weight //	연결이 가능하다면 연결
			cnt++
		}
		if cnt == isLand-2 { //	연결 수행 횟수가 정점-1개가 되면 MST 생성 완료
			break
		}
	}

	//	간선의 수가 n-1이 되지 않는다면 MST 생성 불가
	//	이를 만족하는 경우에만 최종 길이 출력
	if cnt < isLand-2 {
		fmt.Print(-1)
	} else {
		fmt.Print(result)
	}
}

//	BFS로 연결된 섬들 번호 지정
func setNumber(x, y, num int) {
	q := make([]pos, 0)
	maap[x][y] = num
	q = append(q, pos{x, y})
	check[x][y] = true

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		for d := 0; d < 4; d++ {
			nx := cur.x + dx[d]
			ny := cur.y + dy[d]

			if checkRange(nx, ny) {
				continue
			}
			if check[nx][ny] {
				continue
			}
			if maap[nx][ny] == 1 {
				maap[nx][ny] = num
				q = append(q, pos{nx, ny})
				check[nx][ny] = true
			}
		}
	}
}

//	다리 만들기, DFS
func makeBridge(x, y, d, dis int) {
	if checkRange(x, y) { //	범위 넘어가는지 체크
		isEnd = true
		return
	}
	//	1보다 크면 섬이다??
	if maap[x][y] >= 1 {
		//	시삭점과 끝점의 섬 번호가 다르고, 다리 길이가 2 이상인 경우 건설 가능
		if maap[x][y] != num && dis-1 >= 2 {
			//	출발지점, 도착지점, 다리길이 저장
			edges = append(edges, edge{num, maap[x][y], dis - 1})
		}
		//	1보다 큰 경우는 다리 건설이 끝나는 경우
		isEnd = true

		return
	}
	//	이전 진행 방향으로 계속 다리 건설
	makeBridge(x+dx[d], y+dy[d], d, dis+1)

	if isEnd {
		return
	}
}
func union(x, y int) bool {
	xRoot := find(x) //	정점 x의 부모 노드 반환
	yRoot := find(y) //	정점 y의 부모 노드 반환

	//	부모가 같지 않다면 연결 가능
	if xRoot != yRoot {
		parents[yRoot] = xRoot
		return true
	}
	return false
}
func find(x int) int {
	if parents[x] < 0 { //	0보다 작은 경우 부모므로 x 반환
		return x
	}
	parents[x] = find(parents[x]) //	아닐경우 부모노드를 찾는다
	return parents[x]
}
func checkRange(x, y int) bool {
	return x < 0 || x >= n || y < 0 || y >= m
}

type pos struct {
	x, y int
}
type edge struct {
	node1, node2, weight int
}

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}
func init() {
	sc.Split(bufio.ScanWords)
}
