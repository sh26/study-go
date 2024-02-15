package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q25314 is a function to solve the problem "코딩은 체육과목 입니다".
// source: https://www.acmicpc.net/problem/25314
func q25314() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	r.Scan()
	num, _ := strconv.Atoi(r.Text())
	defer w.Flush()

	longCount := num / 4
	result := "long "
	for i := 1; i < longCount; i++ {
		result += "long "
	}
	result += "int"
	fmt.Fprintln(w, result)
}
