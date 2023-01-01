// https://www.acmicpc.net/problem/1330

package step04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func q01330() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	if a, b := scanInt(), scanInt(); a < b {
		fmt.Fprint(w, "<")
	} else if a > b {
		fmt.Fprint(w, ">")
	} else {
		fmt.Fprint(w, "==")
	}
}
