// https://school.programmers.co.kr/learn/courses/30/lessons/12937
// 짝수와 홀수
// # 문제
// 정수 num이 짝수일 경우 "Even"을 반환하고 홀수인 경우 "Odd"를 반환하는 함수, solution을 완성해주세요.

// # 제한
// num은 int 범위의 정수입니다.
// 0은 짝수입니다.

package level1

import "fmt"

func q12397() {

	solution := func(num int) string {
		if (num & 1) != 0 {
			return "Odd"
		}
		return "Even"
	}

	fmt.Println(solution(3)) // "Odd"
	fmt.Println(solution(4)) // "Even"
}
