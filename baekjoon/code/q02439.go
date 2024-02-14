package code

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

// q02439 is a function to solve the problem "별 찍기 - 2".
// source: https://www.acmicpc.net/problem/2439
func q02439() {
	const s byte = ' '

	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	for h, c := 1, scanInt(); h <= c; h++ {
		fmt.Fprint(w, string(bytes.Repeat([]byte{s}, c-h)))
		for l := 1; l <= h; l++ {
			fmt.Fprint(w, "*")
		}
		fmt.Fprint(w, "\n")
	}
}
