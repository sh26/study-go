package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q10813 is a function to solve the problem "공 바꾸기".
// source: https://www.acmicpc.net/problem/10813
func q10813() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()
	r.Scan()
	N, _ := strconv.Atoi(r.Text())
	r.Scan()
	M, _ := strconv.Atoi(r.Text())

	baskets := make([]int, N)
	for i := range baskets {
		baskets[i] = i + 1
	}

	for i := 0; i < M; i++ {
		r.Scan()
		from, _ := strconv.Atoi(r.Text())
		r.Scan()
		to, _ := strconv.Atoi(r.Text())

		baskets[from-1], baskets[to-1] = baskets[to-1], baskets[from-1]
	}

	for i, num := range baskets {
		if i != 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprint(w, num)
	}
}
