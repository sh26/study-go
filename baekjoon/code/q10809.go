package code

import (
	"bufio"
	"fmt"
	"os"
)

// q10809 is a function to solve the problem "알파벳 찾기".
// source: https://www.acmicpc.net/problem/10809
func q10809() {
	r, w, letters := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout), make([]int, 26)
	r.Split(bufio.ScanLines)
	defer w.Flush()

	for idx := range letters {
		letters[idx] = -1
	}

	r.Scan()
	for i, x := range r.Text() {
		idx := int(x) - 'a'
		if letters[idx] == -1 {
			letters[idx] = i
		}
	}

	for _, letter := range letters {
		fmt.Fprint(w, letter, " ")
	}
}
