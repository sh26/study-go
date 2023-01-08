package level0

import "fmt"

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

// https://school.programmers.co.kr/learn/courses/30/lessons/120815
// 피자 나눠 먹기 (2)
func q120815() {

	solution := func(n int) int {
		return n / func(a, b int) int {
			for a > 0 {
				a, b = b%a, a
			}
			return b
		}(n, 6)
	}

	fmt.Println(solution(6))  // 1
	fmt.Println(solution(10)) // 5
	fmt.Println(solution(4))  // 2
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120816
// 피자 나눠 먹기 (3)
func q120816() {

	solution := func(slice int, n int) int {
		if n%slice != 0 {
			return n/slice + 1
		}
		return n / slice
		//    (n + slice - 1) / slice
	}

	fmt.Println(solution(7, 10)) // 2
	fmt.Println(solution(4, 12)) // 3
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120817
// 배열의 평균값
func q120817() {

	solution := func(numbers ...int) float64 {

		total := 0

		for _, v := range numbers {
			total += v
		}

		return float64(total) / float64(len(numbers))
	}

	fmt.Println(solution(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))              //	5.5
	fmt.Println(solution(89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99)) //	94.0
}
