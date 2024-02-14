package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q05597 is a function to solve the problem "과제 안 내신 분..?".
// source: https://www.acmicpc.net/problem/5597
func q05597() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()
	scanInt := func() int { r.Scan(); n, _ := strconv.Atoi(r.Text()); return n }

	var bitmap uint32

	for i := 0; i < 28; i++ {
		bitmap |= 1 << (scanInt() - 1)
	}

	f, s := -1, -1
	for i := 0; i < 30; i++ {
		if f == -1 && bitmap&(1<<i) == 0 {
			f = i + 1
		} else if f != -1 && bitmap&(1<<i) == 0 {
			s = i + 1
			break
		}
	}

	fmt.Fprintf(w, "%d\n%d", f, s)
}
