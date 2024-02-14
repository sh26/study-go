package code

import (
	"bufio"
	"fmt"
	"os"
)

// q04344 is a function to solve the problem "평균은 넘겠지".
// source: https://www.acmicpc.net/problem/4344
func q04344() {
	idx1, idx2 := 0, 0

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	for fmt.Fscan(r, &idx1); idx1 > 0; idx1-- {
		avg, count := 0, 0
		fmt.Fscan(r, &idx2)
		data := make([]int, idx2)
		for i := 0; i < idx2; i++ {
			fmt.Fscan(r, &data[i])
			avg += data[i]
		}
		avg /= len(data)

		for _, i := range data {
			if avg < i {
				count++
			}
		}

		fmt.Fprintln(w, fmt.Sprintf("%.3f", (float64(count)/float64(len(data)))*100)+"%")
	}
}
