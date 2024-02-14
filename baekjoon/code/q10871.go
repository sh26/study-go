package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q10871 is a function to solve the problem "X보다 작은 수".
// source: https://www.acmicpc.net/problem/10871
func q10871() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()
	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	for n, x := scanInt(), scanInt(); n > 0; n-- {
		if num := scanInt(); x > num {
			fmt.Fprintf(w, "%d ", num)
		}
	}
}
