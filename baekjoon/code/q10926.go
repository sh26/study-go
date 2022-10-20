// https://www.acmicpc.net/problem/10926

package code

import (
	"bufio"
	. "fmt"
	. "os"
)

func q10926() {
	var n string
	r := bufio.NewReader(Stdin)
	w := bufio.NewWriter(Stdout)

	Fscan(r, &n)

	Fprintln(w, n+"??!")
	w.Flush()
}
