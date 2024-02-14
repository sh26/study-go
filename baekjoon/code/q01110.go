package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q01110 is a function to solve the problem "더하기 사이클".
// source: https://www.acmicpc.net/problem/1110
func q01110() {
	r, w, n, l, t := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout), 0, 0, 0
	r.Split(bufio.ScanWords)
	defer w.Flush()

	r.Scan()
	n, _ = strconv.Atoi(r.Text())

	b := []int{(n - n%10) / 10, n % 10}
	t = n + 1
	for n != t {
		t = b[0] + b[1]
		b[0] = b[1]
		b[1] = t % 10

		t = b[0]*10 + b[1]
		l++
	}

	w.WriteString(fmt.Sprintln(l))
}
