// https://www.acmicpc.net/problem/1008
// # 문제
// 두 정수 A와 B를 입력받은 다음, A/B를 출력하는 프로그램을 작성하시오.

// # 입력
// 첫째 줄에 A와 B가 주어진다. (0 < A, B < 10)

// # 출력
// 첫째 줄에 A/B를 출력한다. 실제 정답과 출력값의 절대오차 또는 상대오차가 10-9 이하이면 정답이다.

// # 예제
// [1 3] → [0.33333333333333333333333333333333]
// [4 5] → [0.8]

package step01

import (
	"fmt"
)

func q01008() {
	var i, j float64
	fmt.Scanf("%f %f", &i, &j)
	fmt.Print(i / j)
}