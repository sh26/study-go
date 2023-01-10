package level0

import (
	"fmt"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120826
// 특정 문자 제거하기
func q120826() {

	solution := func(my_string string, letter string) string {
		return strings.Replace(my_string, letter, "", -1)
	}

	fmt.Println(solution("abcdef", "f")) //	"abcde"
	fmt.Println(solution("BCBdbe", "B")) //	"Cdbe
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120829
// 각도기
func q120829() {

	solution := func(angle int) int {
		if 0 < angle && angle < 90 {
			return 1
		} else if angle == 90 {
			return 2
		} else if 90 < angle && angle < 180 {
			return 3
		} else {
			return 4
		}
	}

	fmt.Println(solution(70))  // 1
	fmt.Println(solution(91))  // 3
	fmt.Println(solution(180)) // 4
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120830
// 양꼬치
func q120830() {

	solution := func(n int, k int) int {
		return 12000*n + 2000*(k-(n/10))
	}

	fmt.Println(solution(10, 3)) //	124,000
	fmt.Println(solution(64, 6)) //	768,000
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120831
// 짝수의 합
func q120831() {

	solution := func(n int) (total int) {
		for i := 0; i <= n; i += 2 {
			total += i
		}
		return
	}

	fmt.Println(solution(10)) // 30
	fmt.Println(solution(4))  // 6
}
