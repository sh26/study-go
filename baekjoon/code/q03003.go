package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// q03003 is a function to solve the problem "킹, 퀸, 룩, 비숍, 나이트, 폰".
// source: https://www.acmicpc.net/problem/3003
func q03003() {
	r, w, must := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout), []int{1, 1, 2, 2, 2, 8}
	defer w.Flush()

	pieces, _ := r.ReadString('\n')
	for idx, piece := range strings.Split(strings.TrimSpace(pieces), " ") {
		count, _ := strconv.Atoi(piece)
		fmt.Fprintf(w, "%d ", must[idx]-count)
	}
}
