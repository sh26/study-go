// https://www.acmicpc.net/problem/10998
// # 문제
// 두 정수 A와 B를 입력받은 다음, A×B를 출력하는 프로그램을 작성하시오.

// # 입력
// 첫째 줄에 A와 B가 주어진다. (0 < A, B < 10)

// # 출력
// 첫째 줄에 A×B를 출력한다.

package code

import (
	"fmt"
)

func q10998() {
	var i, j int
	fmt.Scanf("%d %d", &i, &j)
	fmt.Print(i * j)
}
