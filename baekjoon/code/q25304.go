package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q25304 is a function to solve the problem "영수증".
// source: https://www.acmicpc.net/problem/25304
func q25304() {
	r, w, sum := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout), 0
	r.Split(bufio.ScanWords)
	defer w.Flush()
	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	sum = scanInt()

	for i := scanInt(); i > 0; i-- {
		sum -= (scanInt() * scanInt())
	}

	if sum == 0 {
		fmt.Fprintln(w, "Yes")
	} else {
		fmt.Fprintln(w, "No")
	}
}
