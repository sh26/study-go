// https://www.acmicpc.net/problem/10952
// # 문제
// 두 정수 A와 B를 입력받은 다음, A+B를 출력하는 프로그램을 작성하시오.

// # 입력
// 입력은 여러 개의 테스트 케이스로 이루어져 있다.
// 각 테스트 케이스는 한 줄로 이루어져 있으며, 각 줄에 A와 B가 주어진다. (0 < A, B < 10)
// 입력의 마지막에는 0 두 개가 들어온다.

// # 출력
// 각 테스트 케이스마다 A+B를 출력한다.

// # 예제
// [1 1],[2 3],[3 4],[9 8],[5 2],[0 0] → [2],[5],[7],[17],[7]

package step03

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func q10952() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	for {
		if sum := scanInt() + scanInt(); sum == 0 {
			return
		} else {
			fmt.Fprintln(w, sum)
		}
	}
}
