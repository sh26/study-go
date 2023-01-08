package level0

import (
	"fmt"
	"sort"
)

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
