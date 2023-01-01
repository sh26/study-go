// https://www.acmicpc.net/problem/9498

package step04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func q09498() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	fmt.Fprint(w, string("FFFFFFDCBAA"[scanInt()/10]))
}
