// https://www.acmicpc.net/problem/2439
// # 문제
// 첫째 줄에는 별 1개, 둘째 줄에는 별 2개, N번째 줄에는 별 N개를 찍는 문제

// # 입력
// 첫째 줄에 N(1 ≤ N ≤ 100)이 주어진다.

// # 출력
// 첫째 줄부터 N번째 줄까지 차례대로 별을 출력한다.

// # 예제
// [5] →  [    *],[   **],[  ***],[ ****],[*****]

package code

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

const s byte = ' '

func q02439() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	for h, c := 1, scanInt(); h <= c; h++ {
		fmt.Fprint(w, string(bytes.Repeat([]byte{s}, c-h)))
		for l := 1; l <= h; l++ {
			fmt.Fprint(w, "*")
		}
		fmt.Fprint(w, "\n")
	}
}
