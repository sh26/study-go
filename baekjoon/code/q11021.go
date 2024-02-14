package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q11021 is a function to solve the problem "A+B - 7".
// source: https://www.acmicpc.net/problem/11021
func q11021() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	for a, c := 1, scanInt(); a <= c; a++ {
		fmt.Fprintf(w, "Case #%d: %d\n", a, scanInt()+scanInt())
	}
}
