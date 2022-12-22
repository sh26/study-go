package code

import (
	"fmt"
	"sort"
)

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
