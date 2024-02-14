package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q02525 is a function to solve the problem "오븐 시계".
// source: https://www.acmicpc.net/problem/2525
func q02525() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	h, m := scanInt(), scanInt()
	if m += h*60 + scanInt(); m >= 1440 {
		m -= 1440
	}
	fmt.Fprintln(w, m/60, m%60)
}
