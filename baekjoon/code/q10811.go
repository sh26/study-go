package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q10811 is a function to solve the problem "바구니 뒤집기".
// source: https://www.acmicpc.net/problem/10811
func q10811() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	r.Scan()
	N, _ := strconv.Atoi(r.Text())
	r.Scan()
	M, _ := strconv.Atoi(r.Text())

	reverseRanges := make([][2]int, M)
	for i := range reverseRanges {
		r.Scan()
		start, _ := strconv.Atoi(r.Text())
		r.Scan()
		end, _ := strconv.Atoi(r.Text())
		reverseRanges[i] = [2]int{start, end}
	}

	baskets := make([]int, N)
	for i := range baskets {
		baskets[i] = i + 1
	}

	for _, r := range reverseRanges {
		start, end := r[0]-1, r[1]-1
		for start < end {
			baskets[start], baskets[end] = baskets[end], baskets[start]
			start++
			end--
		}
	}

	for i, num := range baskets {
		if i != 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprint(w, num)
	}
}
