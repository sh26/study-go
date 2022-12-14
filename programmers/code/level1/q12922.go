// https://school.programmers.co.kr/learn/courses/30/lessons/12922
// 수박수박수박수박수박수?
// # 문제
// 길이가 n이고, "수박수박수박수...."와 같은 패턴을 유지하는 문자열을 리턴하는 함수, solution을 완성하세요.
// 예를들어 n이 4이면 "수박수박"을 리턴하고 3이라면 "수박수"를 리턴하면 됩니다.

// # 제한
// n은 길이 10,000이하인 자연수입니다.

package level1

import (
	"fmt"
	"strings"
)

func q12922() {

	solution := func(n int) string {
		return strings.Repeat("수박", n/2) + strings.Repeat("수", n%2)
	}

	fmt.Println(solution(3)) // "수박수"
	fmt.Println(solution(4)) // "수박수박"
}
