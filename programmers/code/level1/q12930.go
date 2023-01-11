// https://school.programmers.co.kr/learn/courses/30/lessons/12930
// 이상한 문자 만들기
// # 문제
// 문자열 s는 한 개 이상의 단어로 구성되어 있습니다.
// 각 단어는 하나 이상의 공백문자로 구분되어 있습니다.
// 각 단어의 짝수번째 알파벳은 대문자로, 홀수번째 알파벳은 소문자로 바꾼 문자열을 리턴하는 함수, solution을 완성하세요.

// # 제한
// 문자열 전체의 짝/홀수 인덱스가 아니라, 단어(공백을 기준)별로 짝/홀수 인덱스를 판단해야합니다.
// 첫 번째 글자는 0번째 인덱스로 보아 짝수번째 알파벳으로 처리해야 합니다.
package level1

import (
	"bytes"
	"fmt"
	"strings"
)

func q12930() {

	solution := func(s string) string {
		var b bytes.Buffer
		for i, flip, max := 0, false, len(s); i < max; i, flip = i+1, !flip {
			switch {
			case s[i] == ' ':
				flip = true
				fmt.Fprint(&b, " ")
			case flip:
				fmt.Fprint(&b, strings.ToLower(s[i:i+1]))
			default:
				fmt.Fprint(&b, strings.ToUpper(s[i:i+1]))
			}
		}

		return b.String()
	}

	fmt.Println(solution("try hello world")) // "TrY HeLlO WoRlD"
}
