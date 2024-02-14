package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q14681 is a function to solve the problem "사분면 고르기".
// source: https://www.acmicpc.net/problem/14681
func q14681() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	if x, y := scanInt(), scanInt(); x > 0 && y > 0 {
		fmt.Fprintln(w, "1")
	} else if x > 0 && y < 0 {
		fmt.Fprintln(w, "4")
	} else if x < 0 && y > 0 {
		fmt.Fprintln(w, "2")
	} else {
		fmt.Fprintln(w, "3")
	}
}
