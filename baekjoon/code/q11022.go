package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q11022 is a function to solve the problem "A+B - 8".
// source: https://www.acmicpc.net/problem/11022
func q11022() {
	r, w, num := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout), map[byte]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() int {
		r.Scan()
		return num[r.Bytes()[0]]
	}

	for a, c := 1, func() int { r.Scan(); res, _ := strconv.Atoi(r.Text()); return res }(); a <= c; a++ {
		n1, n2 := scanInt(), scanInt()
		fmt.Fprintf(w, "Case #%d: %d + %d = %d\n", a, n1, n2, n1+n2)
	}
}
