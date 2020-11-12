package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	br  = bufio.NewReader(os.Stdin)
	bw  = bufio.NewWriter(os.Stdout)
	s   string
	str []string
	st  = make([]string, 0)
)

func main() {
	defer bw.Flush()
	fmt.Fscan(br, &s)
	if len(s)%2 == 1 { //	짝수 아니면 올바른 괄호가 아님
		fmt.Print(0)
	} else {
		str = strings.Split(s, "")
		fmt.Fprint(bw, find())
	}
}
func find() int {
	//	괄호를 다 풀어서 계산
	//	ex) ([[]()[]])
	//	2*(3*(3+2+3))
	//	2*3*3 + 2*3*2 + 2*3*3

	result, cur := 0, 1 //	결과, 현재 곱할 값
	for i := 0; i < len(str); i++ {
		if str[i] == "(" { //	여는 괄호엔 2를 곱한 후 push
			cur *= 2
			st = append(st, str[i])
		} else if str[i] == "[" { //	여는 괄호엔 3을 곱한 후 push
			cur *= 3
			st = append(st, str[i])
		} else if str[i] == ")" {
			if len(st) == 0 || st[len(st)-1] != "(" { //	스택 길이가 0이거나 여는 괄호가 아니면 종료
				result = 0
				break
			}
			if str[i-1] == "(" { //	올바른 괄호는 더해주고 pop
				result += cur
			}
			st = st[:len(st)-1]
			cur /= 2
		} else if str[i] == "]" {
			if len(st) == 0 || st[len(st)-1] != "[" { //	스택 길이가 0이거나 여는 괄호가 아니면 종료
				result = 0
				break
			}
			if str[i-1] == "[" { //	올바른 괄호는 더해주고 pop
				result += cur
			}
			st = st[:len(st)-1]
			cur /= 3
		}
	}
	return result
}
