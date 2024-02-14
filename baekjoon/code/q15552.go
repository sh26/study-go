package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q15552 is a function to solve the problem "빠른 A+B".
// source: https://www.acmicpc.net/problem/15552
func q15552() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	for c := scanInt(); c > 0; c-- {
		fmt.Fprintln(w, scanInt()+scanInt())
	}
}
