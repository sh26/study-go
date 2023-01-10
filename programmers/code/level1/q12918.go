// https://school.programmers.co.kr/learn/courses/30/lessons/12918
// 문자열 다루기 기본
// # 문제
// 문자열 s의 길이가 4 혹은 6이고, 숫자로만 구성돼있는지 확인해주는 함수, solution을 완성하세요.
// 예를 들어 s가 "a234"이면 False를 리턴하고 "1234"라면 True를 리턴하면 됩니다.

// # 제한
// s는 길이 1 이상, 길이 8 이하인 문자열입니다.
// s는 영문 알파벳 대소문자 또는 0부터 9까지 숫자로 이루어져 있습니다.

package level1

import (
	"fmt"
	"strings"
)

func q12918() {

	solution := func(s string) bool {
		return (len(s) == 6 || len(s) == 4) && len(strings.NewReplacer("0", "", "1", "", "2", "", "3", "", "4", "", "5", "", "6", "", "7", "", "8", "", "9", "").Replace(s)) == 0
		// return (len(s) == 4 || len(s) ==  6) && strings.ToUpper(s) == s && strings.ToLower(s) == s
	}

	fmt.Println(solution("a234")) // false
	fmt.Println(solution("1234")) // true
}
