package code

import (
	"bufio"
	"fmt"
	"os"
)

// q18108 is a function to solve the problem "1998년생인 내가 태국에서는 2541년생?!".
// source: https://www.acmicpc.net/problem/18108
func q18108() {
	r, w, n := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout), 0
	defer w.Flush()
	fmt.Fscan(r, &n)

	fmt.Fprintln(w, n-543)
}
