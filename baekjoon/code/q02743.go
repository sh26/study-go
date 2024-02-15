package code

import (
	"bufio"
	"fmt"
	"os"
)

// q02743 is a function to solve the problem "단어 길이 재기".
// source: https://www.acmicpc.net/problem/2743
func q02743() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()
	r.Scan()
	fmt.Fprint(w, len(r.Text()))
}
