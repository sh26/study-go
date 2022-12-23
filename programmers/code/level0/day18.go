package code

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120908
// 문자열안에 문자열
func q120908() {

	solution := func(str1 string, str2 string) (val int) {
		defer func() { /* :D */
			if err := recover(); err != nil {
				val = 2
			}
		}()
		return strings.Count(str1, str2) / strings.Count(str1, str2)
		// if strings.Count(str1, str2) > 0 {
		// 	return 1
		// }
		// return 0
	}

	fmt.Println(solution("ab6CDE443fgh22iJKlmn1o", "6CD")) // 1
	fmt.Println(solution("ppprrrogrammers", "pppp"))       // 2
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120909
// 제곱수 판별하기
func q120909() {

	solution := func(n int) int {
		if float64(int(math.Sqrt(float64(n)))) < math.Sqrt(float64(n)) {
			return 2
		}
		return 1
	}

	fmt.Println(solution(144)) // 1
	fmt.Println(solution(976)) // 2
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120910
// 세균 증식
func q120910() {

	solution := func(n int, t int) int {
		return n * int(math.Pow(2, float64(t)))
	}

	fmt.Println(solution(2, 10)) // 2048
	fmt.Println(solution(7, 15)) // 229,376
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120911
// 문자열 정렬하기 (2)
func q120911() {

	solution := func(my_string string) string {
		str := strings.Split(strings.ToLower(my_string), "")
		sort.Strings(str)
		return strings.Join(str, "")
	}

	fmt.Println(solution("Bcad"))   // "abcd"
	fmt.Println(solution("heLLo"))  // "ehllo"
	fmt.Println(solution("Python")) // "hnopty"
}
