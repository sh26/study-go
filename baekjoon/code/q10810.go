package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q10810 is a function to solve the problem "공 넣기".
// source: https://www.acmicpc.net/problem/10810
func q10810() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()
	r.Scan()
	N, _ := strconv.Atoi(r.Text())
	r.Scan()
	M, _ := strconv.Atoi(r.Text())

	baskets := make([]int, N)
	for i := 0; i < M; i++ {
		r.Scan()
		start, _ := strconv.Atoi(r.Text())
		r.Scan()
		end, _ := strconv.Atoi(r.Text())
		r.Scan()
		num, _ := strconv.Atoi(r.Text())
		for j := start - 1; j <= end-1; j++ {
			baskets[j] = num
		}
	}

	for i, num := range baskets {
		if i != 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprint(w, num)
	}
	w.Flush()
}
