package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	br = bufio.NewReader(os.Stdin)
	bw = bufio.NewWriter(os.Stdout)
	s  []byte
)

func main() {
	fmt.Scan(&s)
	fmt.Println(find())
}
func find() int {
	r := 0
	for {
		switch {
		case len(s) == 0: //	종료
			return r
		case s[0] == ')': //	닫히면 pop
			s = s[1:]
			return r
		case len(s) > 1 && s[1] == '(': //	열리는 괄호면 곱하는 수 찾고, push
			k := int(s[0]) - '0'
			s = s[2:]
			r += k * find() //	현재 문자부터 다시 탐색
		default: //	그 외에는 인덱싱 하면서 count
			s = s[1:]
			r++
		}
	}
}
