package level0

import (
	"fmt"
	"strconv"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120871
// 저주의 숫자 3
func q120871() {

	solution := func(n int) int {
		for dec, cnt := 1, 1; ; dec++ {
			for { // 3의 배수거나 10 이상이며 수에 3이 포함되는 경우
				if cnt%3 == 0 || strings.Contains(strconv.Itoa(cnt), "3") {
					cnt++
				} else {
					if dec >= n {
						return cnt
					}
					cnt++
					break
				}
			}
		}
	}
	fmt.Println(solution(15)) // 25
	fmt.Println(solution(40)) // 76
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120875
// 평행
func q120875() {

	solution := func(dots ...[]int) int {

		slope := func(dot1, dot2 []int) float64 {
			return float64(dot1[1]-dot2[1]) / float64(dot1[0]-dot2[0])
		}

		if slope(dots[0], dots[1]) == slope(dots[2], dots[3]) || slope(dots[0], dots[2]) == slope(dots[1], dots[3]) || slope(dots[0], dots[3]) == slope(dots[1], dots[2]) {
			return 1
		}

		return 0
	}

	fmt.Println(solution([]int{1, 4}, []int{9, 2}, []int{3, 8}, []int{11, 6})) // 1
	fmt.Println(solution([]int{3, 5}, []int{4, 1}, []int{2, 4}, []int{5, 10})) // 0
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120876
// 겹치는 선분의 길이
func q120876() {

	solution := func(lines ...[]int) (cnt int) {

		elements := make([]int, 201)

		stamp := func(element int) {
			elements[element+100]++
		}

		for _, line := range lines {
			for i := line[0]; i < line[1]; i++ {
				stamp(i)
			}
		}

		for _, element := range elements {
			if element >= 2 {
				cnt++
			}
		}

		return
	}

	fmt.Println(solution([]int{0, 1}, []int{2, 5}, []int{3, 9}))  // 2
	fmt.Println(solution([]int{-1, 1}, []int{1, 3}, []int{3, 9})) // 0
	fmt.Println(solution([]int{0, 5}, []int{3, 9}, []int{1, 10})) // 8
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120878
// 유한소수 판별하기
func q120878() {

	solution := func(a int, b int) int {

		gcd := func(n1, n2 int) int {
			for n1 > 0 {
				n1, n2 = n2%n1, n1
			}
			return b / n2
		}(a, b)

		for i := 3; i <= gcd; i += 2 {
			if i%5 == 0 {
				continue
			}

			if gcd%i == 0 {
				return 2
			}
		}

		return 1
	}

	fmt.Println(solution(7, 20))  // 1
	fmt.Println(solution(11, 20)) // 1
	fmt.Println(solution(12, 21)) // 2
}
