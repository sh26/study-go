package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q01330 is a function to solve the problem "두 수 비교하기".
// source: https://www.acmicpc.net/problem/1330
func q01330() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	if a, b := scanInt(), scanInt(); a < b {
		fmt.Fprint(w, "<")
	} else if a > b {
		fmt.Fprint(w, ">")
	} else {
		fmt.Fprint(w, "==")
	}
}
