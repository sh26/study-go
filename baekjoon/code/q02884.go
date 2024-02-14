package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q02884 is a function to solve the problem "알람 시계".
// source: https://www.acmicpc.net/problem/2884
func q02884() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	h, m := scanInt(), scanInt()
	if m += h*60 - 45; m < 0 {
		m += 1440
	}
	fmt.Fprintln(w, m/60, m%60)
}
