// https://school.programmers.co.kr/learn/courses/30/lessons/12934
// 정수 제곱근 판별
// # 문제
// 임의의 양의 정수 n에 대해, n이 어떤 양의 정수 x의 제곱인지 아닌지 판단하려 합니다.
// n이 양의 정수 x의 제곱이라면 x+1의 제곱을 리턴하고, n이 양의 정수 x의 제곱이 아니라면 -1을 리턴하는 함수를 완성하세요.

// # 제한
// n은 1이상, 50000000000000 이하인 양의 정수입니다.

package level1

import (
	"fmt"
	"math"
)

func q12934() {

	solution := func(n int64) int64 {
		if s := math.Sqrt(float64(n)); s == float64(int(s)) {
			return int64(math.Pow(s+1, 2))
		}
		return -1
	}

	fmt.Println(solution(121)) // 144
	fmt.Println(solution(3))   // -1
}
