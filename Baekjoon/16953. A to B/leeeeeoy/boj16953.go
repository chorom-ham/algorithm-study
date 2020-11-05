package main

import (
	"fmt"
	"strconv"
)

var (
	a, b int
	max  = 1000000000
)

//	162 -> 81 -> 8 -> 4 -> 2 -> 1
//	42 -> 21 -> 2 -> 1
//	40021 -> 4002 -> 2001 -> 200 -> 100 -> 50 -> 25

//	b에서 a찾기
func main() {
	fmt.Scan(&a, &b)
	count := 1
	for b != a {
		if b < a { //	b가 a보다 작아지면 못찾음
			count = -1
			break
		}
		bb := strconv.Itoa(b)
		if bb[len(bb)-1] != '1' && b%2 != 0 { //	끝자리가 1이 아닌데 2로 안떨어지면 못찾음
			count = -1
			break
		}

		//	위의 경우를 제외하고 이 2가지 경우 외엔 없음
		if b%2 == 0 { //	2로 나눠질땐 2로 나눠줌
			b /= 2
		} else { //	그 외엔 1을 빼줌
			bb = bb[:len(bb)-1]
			b, _ = strconv.Atoi(bb)
		}
		count++
	}
	fmt.Print(count)
}

//BFS, 통과안됨
func find() int {
	count, cur := 1, 0
	q := make([]int, 0)
	q = append(q, a)
	for len(q) > 0 {
		k := len(q)
		for i := 0; i < k; i++ {
			cur = q[0]
			q = q[1:]
			if cur == b {
				return count
			}
			if cur > max {
				return -1
			}
			q = append(q, cur*2)
			q = append(q, cur*10+1)
		}
		count++
	}
	return -1
}
