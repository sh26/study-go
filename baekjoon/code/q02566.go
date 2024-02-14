package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// q02566 is a function to solve the problem "최댓값".
// source: https://www.acmicpc.net/problem/2566
func q02566() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		max        rune
		pos1, pos2 int
	)

	for i := 0; i < 9; i++ {
		line, _ := r.ReadString('\n')
		for j, word := range strings.Split(strings.TrimSpace(line), " ") {
			if num, _ := strconv.Atoi(word); int(max) < num {
				max, pos1, pos2 = rune(num), i, j
			}
		}
	}
	fmt.Fprintf(w, "%d\n%d %d", max, pos1+1, pos2+1)
}
