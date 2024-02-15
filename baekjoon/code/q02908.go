package code

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// q02908 is a function to solve the problem "상수".
// source: https://www.acmicpc.net/problem/2908
func q02908() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	r.Split(bufio.ScanLines)
	r.Scan()
	str := strings.Split(r.Text(), " ")

	a, b := func() (a []byte, b []byte) {
		a, b = []byte(str[0]), []byte(str[1])
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
		return
	}()
	if string(a) > string(b) {
		fmt.Fprint(w, string(a))
		return
	}
	fmt.Fprint(w, string(b))
}
