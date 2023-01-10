// https://school.programmers.co.kr/learn/courses/30/lessons/12948
// 핸드폰 번호 가리기
// # 문제
// 프로그래머스 모바일은 개인정보 보호를 위해 고지서를 보낼 때 고객들의 전화번호의 일부를 가립니다.
// 전화번호가 문자열 phone_number로 주어졌을 때,
// 전화번호의 뒷 4자리를 제외한 나머지 숫자를 전부 *으로 가린 문자열을 리턴하는 함수, solution을 완성해주세요.

// # 제한
// phone_number는 길이 4 이상, 20이하인 문자열입니다.

package level1

import (
	"fmt"
	"strings"
)

func q12948() {

	solution := func(phone_number string) string {
		return strings.NewReplacer("0", "*", "1", "*", "2", "*", "3", "*", "4", "*", "5", "*", "6", "*", "7", "*", "8", "*", "9", "*").Replace(phone_number[:len(phone_number)-4]) + phone_number[len(phone_number)-4:]
		// return strings.Repeat("*", len(phone_number)-4) + phone_number[len(phone_number)-4:]
	}

	fmt.Println(solution("01033334444")) // *******4444
	fmt.Println(solution("027778888"))   // *****8888
}
