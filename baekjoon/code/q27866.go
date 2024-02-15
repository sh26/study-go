package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q27866 is a function to solve the problem "문자와 문자열".
// source: https://www.acmicpc.net/problem/27866
func q27866() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	r.Scan()
	str := r.Text()
	defer w.Flush()
	fmt.Fprintf(w, "%c\n", str[func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}()-1])
}
