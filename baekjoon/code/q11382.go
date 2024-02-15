package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q11382 is a function to solve the problem "꼬마 정민".
// source: https://www.acmicpc.net/problem/11382
func q11382() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()
	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}
	fmt.Fprint(w, scanInt()+scanInt()+scanInt())
}
