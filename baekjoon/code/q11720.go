package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q11720 is a function to solve the problem "숫자의 합".
// source: https://www.acmicpc.net/problem/11720
func q11720() {
	r, w, sum := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout), 0
	r.Split(bufio.ScanWords)
	defer w.Flush()
	scanInt := func() (num int) { r.Scan(); num, _ = strconv.Atoi(r.Text()); return }
	scanBytes := func() (nums []byte) { r.Scan(); nums = r.Bytes(); return }
	for i, j, nums := 0, scanInt(), scanBytes(); i < j; i++ {
		sum += int(nums[i] - 48)
	}
	fmt.Fprintf(w, "%d\n", sum)
}
