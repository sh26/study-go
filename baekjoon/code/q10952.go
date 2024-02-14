package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q10952 is a function to solve the problem "A+B - 5".
// source: https://www.acmicpc.net/problem/10952
func q10952() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	for {
		if sum := scanInt() + scanInt(); sum == 0 {
			return
		} else {
			fmt.Fprintln(w, sum)
		}
	}
}
