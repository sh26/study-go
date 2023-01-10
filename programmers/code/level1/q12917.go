// https://school.programmers.co.kr/learn/courses/30/lessons/12917
// 문자열 내림차순으로 배치하기
// # 문제
// 문자열 s에 나타나는 문자를 큰것부터 작은 순으로 정렬해 새로운 문자열을 리턴하는 함수, solution을 완성해주세요.
// s는 영문 대소문자로만 구성되어 있으며, 대문자는 소문자보다 작은 것으로 간주합니다.

// # 제한
// str은 길이 1 이상인 문자열입니다.

package level1

import (
	"fmt"
	"sort"
	"strings"
)

func q12917() {

	solution := func(s string) string {
		str := strings.Split(s, "")
		sort.Sort(sort.Reverse(sort.StringSlice(str)))
		return strings.Join(str, "")
	}

	fmt.Println(solution("Zbcdefg")) // "gfedcbZ"
}
