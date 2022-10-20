package code

import (
	"fmt"
	"math"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120835
// 순서쌍의 개수
func q120836() {

	solution := func(n int) int {

		fmt.Println(math.Sqrt(float64(n)))

		return 0
	}

	fmt.Println(solution(20))  // 6
	fmt.Println(solution(100)) // 9
}
