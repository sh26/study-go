package code

import (
	"bufio"
	"fmt"
	"os"
)

// q04673 is a function to solve the problem "셀프 넘버".
// source: https://www.acmicpc.net/problem/4673
func q04673() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	nums := make([]bool, 10000)

	for i := range nums {
		sum := i + 1

		for j := i; j != 0; j /= 10 {
			sum += j % 10
		}

		if sum <= 10000 {
			nums[sum-1] = true
		}
	}

	for i := range nums {
		if !nums[i] {
			fmt.Fprintln(w, i)
		}
	}
}
