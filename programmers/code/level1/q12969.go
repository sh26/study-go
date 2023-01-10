// https://school.programmers.co.kr/learn/courses/30/lessons/12969
// 직사각형 별찍기
// # 문제
// 이 문제에는 표준 입력으로 두 개의 정수 n과 m이 주어집니다.
// 별(*) 문자를 이용해 가로의 길이가 n, 세로의 길이가 m인 직사각형 형태를 출력해보세요.

// # 제한
// n과 m은 각각 1000 이하인 자연수입니다.

// # 예제
// [5 3] → *****
// 		   *****
// 		   *****

package level1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func q12969() {

	solution := func() {
		r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
		r.Split(bufio.ScanWords)
		defer w.Flush()

		scanInt := func() (num int) {
			r.Scan()
			num, _ = strconv.Atoi(r.Text())
			return
		}

		for x, y := scanInt(), scanInt(); y > 0; y-- {
			for n := 0; n < x; n++ {
				fmt.Fprint(w, "*")
			}
			fmt.Fprintln(w, "")
		}
	}

	// scan 5 3
	solution()
}
