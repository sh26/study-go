package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q02675 is a function to solve the problem "문자열 반복".
// source: https://www.acmicpc.net/problem/2675
func q02675() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()
	scanInt := func() (num int) { r.Scan(); num, _ = strconv.Atoi(r.Text()); return }
	scanBytes := func() (bytes []byte) { r.Scan(); bytes = []byte(r.Text()); return }
	for i := scanInt(); i > 0; i-- {
		j := scanInt()
		for _, word := range scanBytes() {
			for k := j; k > 0; k-- {
				fmt.Fprintf(w, "%c", word)
			}
		}
		fmt.Fprintln(w)
	}
}
