package code

import (
	"bufio"
	"fmt"
	"os"
)

// q05622 is a function to solve the problem "다이얼".
// source: https://www.acmicpc.net/problem/5622
func q05622() {
	reader := bufio.NewReader(os.Stdin)
	w, ret := bufio.NewWriter(os.Stdout), 0
	defer w.Flush()

	for {
		c, _ := reader.ReadByte()
		if c == '\n' {
			break
		}

		if 'Z' <= c {
			c--
		}
		if 'S' <= c {
			c--
		}

		ret += int(c-65)/3 + 3
	}

	fmt.Fprint(w, ret)
}
