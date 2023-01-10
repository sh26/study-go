// https://school.programmers.co.kr/learn/courses/30/lessons/12928
// 약수의 합
// # 문제
// 정수 n을 입력받아 n의 약수를 모두 더한 값을 리턴하는 함수, solution을 완성해주세요.

// # 제한
// n은 0 이상 3000이하인 정수입니다.

package level1

import "fmt"

func q12928() {
	solution := func(n int) (sum int) {
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				sum += i
			}
		}
		return
	}

	fmt.Println(solution(12)) // 28
	fmt.Println(solution(5))  // 6
}
