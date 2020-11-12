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

func main() {
	defer bw.Flush()
	var str string
	fmt.Fscan(br, &str)
	count := 0

	//	PPAP
	// PPAPPAP, PPPAPAP, PPAPPAP -> PPAP
	// PPAPAPP -> NP
	if len(str) <= 3 {
		if str == "P" { //	P도 PPAP...
			fmt.Fprint(bw, "PPAP")
		} else {
			fmt.Fprint(bw, "NP")
		}
	} else {
		for i := 0; i < len(str); i++ {
			if str[i] == 'P' {
				count++
				continue
			}
			if i != len(str)-1 && count >= 2 && str[i+1] == 'P' {
				count--
				i++
			} else {
				fmt.Fprint(bw, "NP")
				return
			}
		}

		if count == 1 {
			fmt.Fprint(bw, "PPAP")
		} else {
			fmt.Fprint(bw, "NP")
		}
	}
}
