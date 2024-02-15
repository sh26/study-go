package code

import (
	"bufio"
	"fmt"
	"os"
)

// q11654 is a function to solve the problem "아스키 코드".
// source: https://www.acmicpc.net/problem/11654
func q11654() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()
	r.Scan()
	fmt.Fprint(w, r.Text()[0])
}
