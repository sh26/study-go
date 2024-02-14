package code

import (
	"bufio"
	"fmt"
	"os"
)

// q10951 is a function to solve the problem "A+B - 4".
// source: https://www.acmicpc.net/problem/10951
func q10951() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for r.Scan() {
		n := r.Bytes()
		w.WriteString(fmt.Sprintln(int(n[0]-'0') + int(n[2]-'0')))
	}
}
