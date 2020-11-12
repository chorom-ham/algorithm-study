package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
)

//	많이 느림
func main() {
	defer bw.Flush()
	var n, x int
	fmt.Fscan(br, &n, &x)

	//	범위 초과 체크
	if n > x || x > n*26 {
		fmt.Fprint(bw, "!")
		return
	}
	answer := make([]byte, 0)

	//	처음에 모두 A로 채운다
	for i := 0; i < n; i++ {
		answer = append(answer, 'A')
	}
	x -= n //	A로 채운만큼 빼줌

	//	뒤에서부터 큰값으로 맞춰준다
	for i := n - 1; i >= 0 && x > 0; i-- {
		cur := minInt(x, 25)
		answer[i] += byte(cur)
		x -= cur
	}
	for _, s := range answer {
		fmt.Fprint(bw, string(s))
	}
}
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
