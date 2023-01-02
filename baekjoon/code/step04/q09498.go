// https://www.acmicpc.net/problem/9498
// # 문제
// 시험 점수를 입력받아 90 ~ 100점은 A, 80 ~ 89점은 B, 70 ~ 79점은 C, 60 ~ 69점은 D, 나머지 점수는 F를 출력하는 프로그램을 작성하시오.

// # 입력
// 첫째 줄에 시험 점수가 주어진다. 시험 점수는 0보다 크거나 같고, 100보다 작거나 같은 정수이다.

// # 출력
// 시험 성적을 출력한다.

// # 예제
// [100] → [A]

package step04

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func q09498() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	fmt.Fprint(w, string("FFFFFFDCBAA"[scanInt()/10]))
}
