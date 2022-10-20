package code

import (
	"fmt"
	"sort"
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

// https://school.programmers.co.kr/learn/courses/30/lessons/120810
// 나머지 구하기
func q120810() {

	solution := func(num1 int, num2 int) int {
		return num1 % num2
	}

	fmt.Println(solution(3, 2))  // 1
	fmt.Println(solution(10, 5)) // 0
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120811
// 중앙값 구하기
func q120811() {

	solution := func(array ...int) int {
		sort.Ints(array)
		return array[len(array)/2]
	}

	fmt.Println(solution(1, 2, 7, 10, 11)) // 7
	fmt.Println(solution(9, -1, 0))        // 0
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120812
// 최빈값 구하기
func q120812() {

	solution := func(array ...int) int {

		count := map[int]int{}

		max, high := -1, -1

		for _, v := range array {
			count[v]++
		}

		for i, v := range count {
			if max < v {
				max = v
				high = i
			} else if max == v {
				high = -1
			}
		}

		return high
	}

	fmt.Println(solution(1, 2, 3, 3, 3, 4)) // 3
	fmt.Println(solution(1, 1, 2, 2))       // -1
	fmt.Println(solution(1))                // 1
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120813
// 짝수는 싫어요
func q120813() {

	solution := func(n int) (v []int) {
		for i := 1; i <= n; i += 2 {
			v = append(v, i)
		}

		return
	}

	fmt.Println(solution(10)) //	[1, 3, 5, 7, 9]
	fmt.Println(solution(15)) //	[1, 3, 5, 7, 9, 11, 13, 15]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120814
// 피자 나눠 먹기 (1)
func q120814() {

	solution := func(n int) int {
		if n%7 != 0 {
			return n/7 + 1
		}
		return n / 7
	}

	fmt.Println(solution(7))  // 1
	fmt.Println(solution(1))  // 1
	fmt.Println(solution(15)) // 3
}
