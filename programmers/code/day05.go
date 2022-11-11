package code

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

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

// https://school.programmers.co.kr/learn/courses/30/lessons/120849
// 모음 제거
func q120849() {

	solution := func(my_string string) string {
		return strings.NewReplacer("a", "", "e", "", "i", "", "o", "", "u", "").Replace(my_string)
		//     regexp.MustCompile("[aeiou]").ReplaceAllString(my_string, "")
	}

	fmt.Println(solution("bus"))              // "bs"
	fmt.Println(solution("nice to meet you")) // "nc t mt y"
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120850
// 문자열 정렬하기 (1)
func q120850() {

	solution := func(my_string string) (nums []int) {
		my_string = regexp.MustCompile("[a-z]").ReplaceAllString(my_string, "")
		for _, v := range my_string {
			n, _ := strconv.Atoi(string(v))
			nums = append(nums, n)
		}
		sort.Ints(nums)

		return
	}

	fmt.Println(solution("hi12392"))   // [1, 2, 2, 3, 9]
	fmt.Println(solution("p2o4i8gj2")) // [2, 2, 4, 8]
	fmt.Println(solution("abcde0"))    // [0]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120851
// 숨어있는 숫자의 덧셈 (1)
func q120851() {

	solution := func(my_string string) (sum int) {
		for _, v := range my_string {
			n, _ := strconv.Atoi(string(v))
			sum += n
		}
		return
	}

	fmt.Println(solution("aAb1B2cC34oOp")) // 10
	fmt.Println(solution("1a2b3c4d123"))   // 16
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120852
// 소인수분해
func q120852() {

	solution := func(n int) (num []int) {
		if n&1 == 0 {
			num = append(num, 2)
		}

		for i := 3; i <= n; i += 2 {
			exists := true
			for _, j := range num {
				if i%j == 0 {
					exists = false
					break
				}
			}

			if exists && n%i == 0 {
				num = append(num, i)
			}
		}

		return

		/*
			t := 2

			for n > 1 {
				if n%t == 0 {
					res = append(res, t)
					n/=t
					for n%t == 0 {
						n/=t
					}
				}
				t++
			}
			return
		*/
	}

	fmt.Println(solution(12))  // [2, 3]
	fmt.Println(solution(17))  // [17]
	fmt.Println(solution(420)) // [2, 3, 5, 7]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120853
// 컨트롤 제트
func q120853() {

	solution := func(s string) (sum int) {

		num_list := []int{}

		for _, v := range strings.Split(s, " ") {

			if v == "Z" {
				if len(num_list) <= 1 {
					num_list = []int{}
				} else {
					num_list = num_list[:len(num_list)-1]
				}
			} else {
				n, _ := strconv.Atoi(v)
				num_list = append(num_list, n)
			}
		}

		for _, v := range num_list {
			sum += v
		}

		return
	}

	fmt.Println(solution("1 2 3 Z Z"))     // 4
	fmt.Println(solution("10 20 30 Z 40")) // 100
	fmt.Println(solution("10 Z 20 Z 1"))   // 1
}
