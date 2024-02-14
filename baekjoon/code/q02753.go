package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q02753 is a function to solve the problem "윤년".
// source: https://www.acmicpc.net/problem/2753
func q02753() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	if n := scanInt(); n%4 == 0 && (n%100 != 0 || n%400 == 0) {
		fmt.Fprintln(w, 1)
	} else {
		fmt.Fprintln(w, 0)
	}
}
