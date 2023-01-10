// https://school.programmers.co.kr/learn/courses/30/lessons/70128
// 내적
// # 문제
// 길이가 같은 두 1차원 정수 배열 a, b가 매개변수로 주어집니다. a와 b의 내적을 return 하도록 solution 함수를 완성해주세요.
// 이때, a와 b의 내적은 a[0]*b[0] + a[1]*b[1] + ... + a[n-1]*b[n-1] 입니다. (n은 a, b의 길이)

// # 제한
// a, b의 길이는 1 이상 1,000 이하입니다.
// a, b의 모든 수는 -1,000 이상 1,000 이하입니다.

package level1

import "fmt"

func q70128() {

	solution := func(a []int, b []int) (sum int) {
		for idx := range a {
			sum += a[idx] * b[idx]
		}
		return
	}

	fmt.Println(solution([]int{1, 2, 3, 4}, []int{-3, -1, 0, 2})) // 3
	fmt.Println(solution([]int{-1, 0, 1}, []int{1, 0, -1}))       // -2
}
