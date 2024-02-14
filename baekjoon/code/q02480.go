package code

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// q02480 is a function to solve the problem "주사위 세개".
// source: https://www.acmicpc.net/problem/2480
func q02480() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	n := []int{scanInt(), scanInt(), scanInt()}

	sort.Ints(n)

	if n[0] == n[2] {
		fmt.Fprintln(w, 10000+(n[0]*1000))
	} else if n[0] == n[1] || n[1] == n[2] {
		fmt.Fprintln(w, 1000+(n[1]*100))
	} else {
		fmt.Fprintln(w, n[2]*100)
	}
}
