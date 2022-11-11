package code

import (
	"fmt"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120841
// 점의 위치 구하기
func q120841() {

	solution := func(dot ...int) int {
		if dot[0] > 0 {
			if dot[1] > 0 {
				return 1
			}
			return 4
		}

		if dot[1] < 0 {
			return 3
		}
		return 2
	}

	fmt.Println(solution(2, 4))  //	1
	fmt.Println(solution(-7, 9)) //	2
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120842
// 2차원으로 만들기
func q120842() {

	solution := func(num_list []int, n int) (num_split [][]int) {

		for i := 0; i < len(num_list); i += n {
			num_split = append(num_split, func() (list []int) {
				for j := i; j < i+n; j++ {
					list = append(list, num_list[j])
				}

				return
			}())
		}

		return
	}

	fmt.Println(solution([]int{1, 2, 3, 4, 5, 6, 7, 8}, 2))           // [1, 2], [3, 4], [5, 6], [7, 8]
	fmt.Println(solution([]int{100, 95, 2, 4, 5, 6, 18, 33, 948}, 3)) // [100, 95, 2], [4, 5, 6], [18, 33, 948]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120843
// 공 던지기
func q120843() {

	solution := func(numbers []int, k int) int {
		return numbers[(2*(k-1))%len(numbers)]
	}

	fmt.Println(solution([]int{1, 2, 3, 4}, 2))       //	3
	fmt.Println(solution([]int{1, 2, 3, 4, 5, 6}, 5)) //	3
	fmt.Println(solution([]int{1, 2, 3}, 3))          //	2
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120844
// 배열 회전시키기
func q120844() {

	solution := func(numbers []int, direction string) []int {
		if direction == "left" {
			return append(numbers[1:], numbers[0])
		}
		return append(numbers[len(numbers)-1:], numbers[0:len(numbers)-1]...)
	}

	fmt.Println(solution([]int{1, 2, 3}, "right"))                // [3, 1, 2]
	fmt.Println(solution([]int{4, 455, 6, 4, -1, 45, 6}, "left")) // [455, 6, 4, -1, 45, 6, 4]
}

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
