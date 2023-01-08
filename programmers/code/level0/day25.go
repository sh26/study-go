package level0

import (
	"fmt"
	"sort"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120921
// 문자열 밀기
func q120921() {
	solution := func(A string, B string) int {
		return strings.Index(B+B, A)
	}

	fmt.Println(solution("hello", "hello")) // 1
	fmt.Println(solution("apple", "elppa")) // -1
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120922
// 종이 자르기
func q120922() {

	solution := func(M int, N int) int {
		return M*N - 1
	}

	fmt.Println(solution(2, 2)) // 3
	fmt.Println(solution(2, 5)) // 9
	fmt.Println(solution(1, 1)) // 0
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120923
// 연속된 수의 합
func q120923() {

	solution := func(num int, total int) (num_list []int) {

		center, cnt := total/num, 1

		for i := 1; i <= num/2; i, cnt = i+1, cnt+1 {
			num_list = append(num_list, center+i)
			cnt++
			if cnt < num {
				num_list = append(num_list, center-i)
			}
		}

		num_list = append(num_list, center)

		sort.Ints(num_list)

		return
	}

	fmt.Println(solution(3, 12)) // [3, 4, 5]
	fmt.Println(solution(5, 15)) // [1, 2, 3, 4, 5]
	fmt.Println(solution(4, 14)) // [2, 3, 4, 5]
	fmt.Println(solution(5, 5))  // [-1, 0, 1, 2, 3]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120924
// 다음에 올 숫자
func q120924() {

	solution := func(common []int) int {
		if val := common[1] - common[0]; val == common[2]-common[1] {
			return common[len(common)-1] + val
		}
		return common[len(common)-1] * (common[1] / common[0])
	}

	fmt.Println(solution([]int{1, 2, 3, 4})) // 5
	fmt.Println(solution([]int{2, 4, 8}))    // 16
}
