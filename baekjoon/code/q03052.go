package code

import (
	"bufio"
	"fmt"
	"os"
)

// q03052 is a function to solve the problem "나머지".
// source: https://www.acmicpc.net/problem/3052
func q03052() {
	rem, num, idx := map[string]bool{}, 0, 10

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	for ; idx > 0; idx-- {
		fmt.Fscanln(r, &num)
		rem[fmt.Sprint(num%42)] = true
	}

	fmt.Fprintln(w, len(rem))
}
