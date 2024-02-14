package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q02438 is a function to solve the problem "별 찍기 - 1".
// source: https://www.acmicpc.net/problem/2438
func q02438() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	for h, c := 1, scanInt(); h <= c; h++ {
		for l := 1; l <= h; l++ {
			fmt.Fprint(w, "*")
		}
		fmt.Fprint(w, "\n")
	}
}
