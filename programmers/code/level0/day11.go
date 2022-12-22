package code

import (
	"fmt"
	"sort"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120845
// 주사위의 개수
func q120845() {

	solution := func(box []int, n int) int {
		res := 1

		for _, v := range box {
			res *= int(v / n)
		}

		return res
	}

	fmt.Println(solution([]int{1, 1, 1}, 1))  // 1
	fmt.Println(solution([]int{10, 8, 6}, 3)) // 12
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120846
// 합성수 찾기
func q120846() {

	solution := func(n int) (cnt int) {

		for i := 4; i <= n; i += 2 {
			cnt++
		}

		for i := 3; i <= n; i += 2 {
			for j := 3; j < i; j++ {
				if i%j == 0 {
					cnt++
					break
				}
			}
		}

		return
	}

	fmt.Println(solution(10)) // 5
	fmt.Println(solution(15)) // 8
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120847
// 최댓값 만들기 (1)
func q120847() {

	solution := func(numbers []int) int {
		sort.Ints(numbers)

		return numbers[len(numbers)-1] * numbers[len(numbers)-2]
	}

	fmt.Println(solution([]int{1, 2, 3, 4, 5}))       // 20
	fmt.Println(solution([]int{0, 31, 24, 10, 1, 9})) // 744
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120848
// 팩토리얼
func q120848() {

	solution := func(n int) int {
		var fact, i int
		for fact, i = 1, 1; fact <= n; i++ {
			fact *= i
		}
		return i - 2
	}

	fmt.Println(solution(3628800)) // 10
	fmt.Println(solution(7))       // 3
}
