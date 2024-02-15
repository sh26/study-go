package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q09086 is a function to solve the problem "문자열".
// source: https://www.acmicpc.net/problem/9086
func q09086() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()
	r.Scan()
	for i, j := 0, func() (num int) { num, _ = strconv.Atoi(r.Text()); return }(); i < j; i++ {
		r.Scan()
		str := r.Text()
		fmt.Fprintf(w, "%c%c\n", str[0], str[len(str)-1])
	}
}
