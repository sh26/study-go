package code

import (
	"bufio"
	"fmt"
	"os"
)

// q10926 is a function to solve the problem "??!".
// source: https://www.acmicpc.net/problem/10926
func q10926() {
	r, w, n := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout), ""
	defer w.Flush()
	fmt.Fscan(r, &n)

	fmt.Fprintln(w, n+"??!")
}
