package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q01065 is a function to solve the problem "한수".
// source: https://www.acmicpc.net/problem/1065
func q01065() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	if limit := scanInt(); limit < 100 {
		fmt.Fprintln(w, limit)
	} else {
		cnt := 99
		for ; limit >= 100; limit-- {
			n1, n2, n3 := limit%10, limit/10%10, limit/100
			if n3-n2 == n2-n1 {
				cnt++
			}
		}
		fmt.Fprintln(w, cnt)
	}
}
