package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// q02745 is a function to solve the problem "진법 변환".
// source: https://www.acmicpc.net/problem/2745
func q02745() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	// 대상
	r.Scan()
	str := r.Bytes()

	// 진법
	r.Scan()
	sys, _ := strconv.Atoi(r.Text())

	// 10진법 값
	num := 0

	for i, j := len(str)-1, 0; i >= 0; i, j = i-1, j+1 {
		x := str[i] - 48
		if x > 9 {
			x -= 7
		}
		num += int(x) * FastExp(sys, j)
	}
	fmt.Fprint(w, num)
}

// 분할 정복 알고리즘 (거듭제곱)
func FastExp(base int, exponent int) int {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	if exponent < 0 {
		return 1 / FastExp(base, -exponent)
	}
	if exponent%2 == 0 {
		half := FastExp(base, exponent/2)
		return half * half
	}
	return base * FastExp(base, exponent-1)
}
