package code

import (
	"fmt"
	"strconv"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120904
// 숫자 찾기
func q120904() {

	solution := func(num int, k int) int {
		return strings.Index(" "+fmt.Sprint(num), fmt.Sprint(k))
	}

	fmt.Println(solution(29183, 1))  // 3
	fmt.Println(solution(232443, 4)) // 4
	fmt.Println(solution(123456, 7)) // -1
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120905
// n의 배수 고르기
func q120905() {

	solution := func(n int, numlist []int) []int {
		nums := []int{}
		for _, num := range numlist {
			if num%n == 0 {
				nums = append(nums, num)
			}
		}

		return nums
	}

	fmt.Println(solution(3, []int{4, 5, 6, 7, 8, 9, 10, 11, 12})) // [6, 9, 12]
	fmt.Println(solution(5, []int{1, 9, 3, 10, 13, 5}))           // [10, 5]
	fmt.Println(solution(12, []int{2, 100, 120, 600, 12, 12}))    // [120, 600, 12, 12]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120906
// 자릿수 더하기
func q120906() {

	solution := func(n int) (res int) {
		for _, val := range strings.Split(fmt.Sprint(n), "") {
			num, _ := strconv.Atoi(val)
			res += num
		}
		return
	}

	fmt.Println(solution(1234))   // 10
	fmt.Println(solution(930211)) // 16
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120907
// OX퀴즈
func q120907() {

	solution := func(quiz []string) (cns []string) {
		for _, full := range quiz {
			expr, res, sum := func() (string, string, int) { expr := strings.Split(full, " = "); return expr[0], expr[1], 0 }()
			for _, val := range strings.Split(strings.NewReplacer("+ ", "+", "- ", "-").Replace("+ "+expr), " ") {
				num, _ := strconv.Atoi(string(val[1:]))
				if string(val[0]) == "+" {
					sum += num
				} else {
					sum -= num
				}
			}

			if fmt.Sprint(sum) == res {
				cns = append(cns, "O")
			} else {
				cns = append(cns, "X")
			}
		}

		return
	}

	fmt.Println(solution([]string{"3 - 4 = -3", "5 + 6 = 11"}))                               // ["X", "O"]
	fmt.Println(solution([]string{"19 - 6 = 13", "5 + 66 = 71", "5 - 15 = 63", "3 - 1 = 2"})) // ["O", "O", "X", "O"]
}
