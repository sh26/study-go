package code

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

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
