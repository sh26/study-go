// https://school.programmers.co.kr/learn/courses/30/lessons/12926
// 시저 암호
// # 문제
// 어떤 문장의 각 알파벳을 일정한 거리만큼 밀어서 다른 알파벳으로 바꾸는 암호화 방식을 시저 암호라고 합니다.
// 예를 들어 "AB"는 1만큼 밀면 "BC"가 되고, 3만큼 밀면 "DE"가 됩니다. "z"는 1만큼 밀면 "a"가 됩니다.
// 문자열 s와 거리 n을 입력받아 s를 n만큼 민 암호문을 만드는 함수, solution을 완성해 보세요.

// # 제한
// 	공백은 아무리 밀어도 공백입니다.
// 	s는 알파벳 소문자, 대문자, 공백으로만 이루어져 있습니다.
// 	s의 길이는 8000이하입니다.
// 	n은 1 이상, 25이하인 자연수입니다.

package level1

import (
	"bytes"
	"fmt"
)

func q12926() {

	solution := func(s string, n int) string {
		var b bytes.Buffer
		for i, max := 0, len(s); i < max; i++ {
			// case 64 <= x && x <= 89 || 96 <= x && x <= 121: // A - Y
			if x := s[i]; x == ' ' {
				fmt.Fprint(&b, " ")
			} else {
				fmt.Fprint(&b, string(x+byte(n)))
			}
		}
		return b.String()
	}

	fmt.Println(solution("AB", 1))                         // "BC"
	fmt.Println(solution("z", 1))                          // "a"
	fmt.Println(solution("aBcDeFgHiJkLmNoPqRsTuVwXyZ", 1)) // "e F d"
	fmt.Println(solution("z Z z", 5))                      // "e F d"
}
