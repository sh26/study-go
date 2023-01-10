// https://school.programmers.co.kr/learn/courses/30/lessons/77884
// 약수의 개수와 덧셈
// # 문제
// 두 정수 left와 right가 매개변수로 주어집니다.
// left부터 right까지의 모든 수들 중에서, 약수의 개수가 짝수인 수는 더하고,
// 약수의 개수가 홀수인 수는 뺀 수를 return 하도록 solution 함수를 완성해주세요.

// # 제한
// 1 ≤ left ≤ right ≤ 1,000

package level1

import "fmt"

func q77884() {

	solution := func(left int, right int) int {
		for sum, cnt := 0, 0; ; left, cnt = left+1, 0 { // 초기화
			for n, m := 1, left; n <= m; n++ { // 약수
				if m%n == 0 {
					cnt++
				}
			}
			if (cnt & 1) != 0 { // 홀수 짝수 판별
				sum -= left
			} else {
				sum += left
			}
			if left > right-1 { // return
				return sum
			}
		}
	}

	fmt.Println(solution(13, 17)) // 43
	fmt.Println(solution(24, 27)) // 52
}
