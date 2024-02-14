package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q10818 is a function to solve the problem "최소, 최대".
// source: https://www.acmicpc.net/problem/10818
func q10818() {
	r, w, min, max := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout), 1000001, -1000001
	r.Split(bufio.ScanWords)
	defer w.Flush()
	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	for cnt := scanInt(); cnt > 0; cnt-- {
		num := scanInt()
		if max < num {
			max = num
		}
		if min > num {
			min = num
		}
	}

	fmt.Fprint(w, min, max)
}
