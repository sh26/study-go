package code

import (
	"bufio"
	"fmt"
	"os"
)

// q02562 is a function to solve the problem "최댓값".
// source: https://www.acmicpc.net/problem/2562
func q02562() {
	size, num, max, idx := 9, 0, -1000001, 0

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	for ; size > 0; size-- {
		fmt.Fscan(r, &num)

		if max < num {
			max = num
			idx = 10 - size
		}
	}
	fmt.Fprintln(w, max)
	fmt.Fprintln(w, idx)
}
