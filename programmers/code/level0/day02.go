package code

import "fmt"

// https://school.programmers.co.kr/learn/courses/30/lessons/120806
// 두 수의 나눗셈
func q120806() {

	solution := func(num1 int, num2 int) int {
		return num1 * 1000 / num2
	}

	fmt.Println(solution(3, 2))  // 1500
	fmt.Println(solution(7, 3))  // 2300
	fmt.Println(solution(1, 16)) // 62
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120807
// 숫자 비교하기
func q120807() {

	solution := func(num1 int, num2 int) int {
		return func() int {
			if num1 == num2 {
				return 1
			}
			return -1
		}()
	}

	fmt.Println(solution(2, 3))   // -1
	fmt.Println(solution(11, 11)) // 1
	fmt.Println(solution(7, 99))  // -1
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120808
// 분수의 덧셈
func q120808() {

	solution := func(denum1 int, num1 int, denum2 int, num2 int) []int {
		val1 := (denum1 * num2) + (denum2 * num1)
		val2 := num1 * num2

		gcd := func(a, b int) int {
			for a > 0 {
				a, b = b%a, a
			}
			return b
		}(val1, val2)

		return []int{val1 / gcd, val2 / gcd}
	}

	fmt.Println(solution(1, 2, 3, 4)) // [5, 4]
	fmt.Println(solution(9, 2, 1, 3)) // [29, 6]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120809
// 배열 두 배 만들기
func q120809() {

	solution := func(numbers ...int) []int {
		for i, v := range numbers {
			numbers[i] = v * 2
		}
		return numbers
	}

	fmt.Println(solution(1, 2, 3, 4, 5))           // [2, 4, 6, 8, 10]
	fmt.Println(solution(1, 2, 100, -99, 1, 2, 3)) // [2, 4, 200, -198, 2, 4, 6]
}
