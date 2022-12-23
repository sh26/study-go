package code

import (
	"fmt"
	"strconv"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120898
// 편지
func q120898() {

	solution := func(message string) int {
		return len(message) * 2
	}

	fmt.Println(solution("happy birthday!")) //	30
	fmt.Println(solution("I love you~"))     //	22
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120899
// 가장 큰 수 찾기
func q120899() {

	solution := func(array []int) []int {
		max, maxidx := -1, -1

		for idx, val := range array {
			if val > max {
				max, maxidx = val, idx
			}
		}

		return []int{max, maxidx}
	}

	fmt.Println(solution([]int{1, 8, 3}))      // [8, 1]
	fmt.Println(solution([]int{9, 10, 11, 8})) // [11, 2]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120902
// 문자열 계산하기
func q120902() {

	solution := func(my_string string) (res int) {
		for _, val := range strings.Split(strings.NewReplacer("+ ", "+", "- ", "-").Replace("+ "+my_string), " ") {
			num, _ := strconv.Atoi(string(val[1:]))
			if string(val[0]) == "+" {
				res += num
			} else {
				res -= num
			}
		}

		return
	}

	fmt.Println(solution("3 + 4")) // 7
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120903
// 배열의 유사도
func q120903() {

	solution := func(s1 []string, s2 []string) (cnt int) {

		full := " " + strings.Join(s2, " ") + " "

		for _, str := range s1 {
			if strings.Count(full, " "+str+" ") > 0 {
				cnt++
			}
		}

		return
	}

	fmt.Println(solution([]string{"a", "b", "c"}, []string{"com", "b", "d", "p", "c"})) // 2
	fmt.Println(solution([]string{"n", "omg"}, []string{"m", "dot"}))                   // 0
}
