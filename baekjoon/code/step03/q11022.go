// https://www.acmicpc.net/problem/11022
// # 문제
// 두 정수 A와 B를 입력받은 다음, A+B를 출력하는 프로그램을 작성하시오.

// # 입력
// 첫째 줄에 테스트 케이스의 개수 T가 주어진다.
// 각 테스트 케이스는 한 줄로 이루어져 있으며, 각 줄에 A와 B가 주어진다. (0 < A, B < 10)

// # 출력
// 각 테스트 케이스마다 "Case #x: A + B = C" 형식으로 출력한다. x는 테스트 케이스 번호이고 1부터 시작하며, C는 A+B이다.

// # 예제
// [5],[1 1],[2 3],[3 4],[9 8],[5 2] → [Case #1: 1 + 1 = 2],[Case #2: 2 + 3 = 5],[Case #3: 3 + 4 = 7],[Case #4: 9 + 8 = 17],[Case #5: 5 + 2 = 7]

package code

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func q11022() {
	r, w, num := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout), map[byte]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() int {
		r.Scan()
		return num[r.Bytes()[0]]
	}

	for a, c := 1, func() int { r.Scan(); res, _ := strconv.Atoi(r.Text()); return res }(); a <= c; a++ {
		n1, n2 := scanInt(), scanInt()
		fmt.Fprintf(w, "Case #%d: %d + %d = %d\n", a, n1, n2, n1+n2)
	}
}
