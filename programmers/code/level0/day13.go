package level0

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

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

// https://school.programmers.co.kr/learn/courses/30/lessons/120854
// 배열 원소의 길이
func q120854() {

	solution := func(strlist []string) (strlen []int) {
		for _, v := range strlist {
			strlen = append(strlen, len(v))
		}

		return
	}

	fmt.Println(solution([]string{"We", "are", "the", "world!"})) // [2, 3, 3, 6]
	fmt.Println(solution([]string{"I", "Love", "Programmers."}))  // [1, 4, 12]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120888
// 중복된 문자 제거
func q120888() {

	solution := func(my_string string) (str string) {
		unique_str := map[string]bool{}
		for _, alp := range my_string {
			if val := unique_str[string(alp)]; !val {
				unique_str[string(alp)] = true
				str += string(alp)
			}
		}

		return
	}

	fmt.Println(solution("people"))           //"peol"
	fmt.Println(solution("We are the world")) //"We arthwold"
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120889
// 삼각형의 완성조건 (1)
func q120889() {

	solution := func(sides ...int) int {
		if sort.Ints(sides); sides[0]+sides[1] > sides[2] {
			return 2
		}
		return 1
	}

	fmt.Println(solution(1, 2, 3))      // 2
	fmt.Println(solution(3, 6, 2))      // 2
	fmt.Println(solution(199, 72, 222)) // 1
}
