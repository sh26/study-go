// https://www.acmicpc.net/problem/1330
// # 문제
// 두 정수 A와 B가 주어졌을 때, A와 B를 비교하는 프로그램을 작성하시오.

// # 입력
// 첫째 줄에 A와 B가 주어진다. A와 B는 공백 한 칸으로 구분되어져 있다.

// # 출력
// 첫째 줄에 다음 세 가지 중 하나를 출력한다.
// A가 B보다 큰 경우에는 '>'를 출력한다.
// A가 B보다 작은 경우에는 '<'를 출력한다.
// A와 B가 같은 경우에는 '=='를 출력한다.
// 제한
// -10,000 ≤ A, B ≤ 10,000

// # 예제
// [1 2] → [<], [10 2] → [>], [5 5] → [==]

package step04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func q01330() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	if a, b := scanInt(), scanInt(); a < b {
		fmt.Fprint(w, "<")
	} else if a > b {
		fmt.Fprint(w, ">")
	} else {
		fmt.Fprint(w, "==")
	}
}
