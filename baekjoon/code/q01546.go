package code

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// q01546 is a function to solve the problem "평균".
// source: https://www.acmicpc.net/problem/1546
func q01546() {
	idx, sum := 0, 0.0

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscan(r, &idx)

	data := make([]float64, idx)

	for i := 0; i < idx; i++ {
		fmt.Fscan(r, &data[i])
		sum += data[i]
	}

	sort.Float64s(data)

	fmt.Fprint(w, sum/float64(idx)*100/data[len(data)-1])
}
