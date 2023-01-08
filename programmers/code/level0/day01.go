package level0

import (
	"fmt"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120802
// 두 수의 합
func q120802() {

	solution := func(num1 int, num2 int) int {
		return num1 + num2
	}

	fmt.Println(solution(2, 3))   // 5
	fmt.Println(solution(100, 2)) // 102
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120803
// 두 수의 차
func q120803() {

	solution := func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Println(solution(2, 3))   // -1
	fmt.Println(solution(100, 2)) // 98
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120804
// 두 수의 곱
func q120804() {

	solution := func(num1 int, num2 int) int {
		return num1 * num2
	}

	fmt.Println(solution(3, 4))   // 12
	fmt.Println(solution(27, 19)) // 513
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120805
// 몫 구하기
func q120805() {

	solution := func(num1 int, num2 int) int {
		return num1 / num2
	}

	fmt.Println(solution(10, 5)) // 2
	fmt.Println(solution(7, 2))  // 3
}
