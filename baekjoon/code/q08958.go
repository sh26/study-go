package code

import (
	"bufio"
	"fmt"
	"os"
)

// q08958 is a function to solve the problem "OX퀴즈".
// source: https://www.acmicpc.net/problem/8958
func q08958() {
	idx := 0

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	text := ""
	for fmt.Fscan(r, &idx); idx > 0; idx-- {
		fmt.Fscan(r, &text)
		sum, suc := 0, 1
		for i := 0; i < len(text); i++ {
			if text[i] == 'O' {
				sum += suc
				suc++
			} else {
				suc = 1
			}
		}
		fmt.Fprintln(w, sum)
	}
}
