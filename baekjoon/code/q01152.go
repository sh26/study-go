package code

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// q01152 is a function to solve the problem "단어의 개수".
// source: https://www.acmicpc.net/problem/1152
func q01152() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	r.Buffer(make([]byte, 1000001), 1000001)
	r.Scan()
	str := bytes.TrimSpace(r.Bytes())
	if len(str) == 0 {
		fmt.Fprint(w, 0)
		return
	}
	fmt.Fprint(w, bytes.Count(str, []byte{' '})+1)
}
