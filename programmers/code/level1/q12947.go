// https://school.programmers.co.kr/learn/courses/30/lessons/12947
// 하샤드 수
// # 문제
// 양의 정수 x가 하샤드 수이려면 x의 자릿수의 합으로 x가 나누어져야 합니다.
// 예를 들어 18의 자릿수 합은 1+8=9이고, 18은 9로 나누어 떨어지므로 18은 하샤드 수입니다.
// 자연수 x를 입력받아 x가 하샤드 수인지 아닌지 검사하는 함수, solution을 완성해주세요.

// # 제한
// x는 1 이상, 10000 이하인 정수입니다.

package level1

import "fmt"

func q12947() {

	solution := func(x int) bool {
		digit := 0
		for n := x; n > 0; digit, n = digit+(n%10), n/10 { }
		return x%digit == 0
	}

	fmt.Println(solution(10)) // true
	fmt.Println(solution(12)) // true
	fmt.Println(solution(11)) // false
	fmt.Println(solution(13)) // false
}
