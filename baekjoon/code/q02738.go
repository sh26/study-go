package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// q02738 is a function to solve the problem "행렬 덧셈".
// source: https://www.acmicpc.net/problem/2738
func q02738() {
	r, w, size := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout), []int{0, 0}
	defer w.Flush()

	fmt.Fscanf(r, "%d %d\n", &size[0], &size[1])

	procs := make([][]int, size[0])

	for i := 0; i < size[0]; i++ {
		procs[i] = make([]int, size[1])
		line, _ := r.ReadString('\n')

		for j, str := range strings.Split(strings.TrimSpace(line), " ") {
			procs[i][j], _ = strconv.Atoi(str)
		}
	}

	for i := 0; i < size[0]; i++ {
		line, _ := r.ReadString('\n')
		for j, str := range strings.Split(strings.TrimSpace(line), " ") {
			num, _ := strconv.Atoi(str)
			procs[i][j] += num
		}
		str := fmt.Sprint(procs[i])

		fmt.Fprintln(w, str[1:len(str)-1])
	}
}
