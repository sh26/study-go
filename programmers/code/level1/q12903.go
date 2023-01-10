// https://school.programmers.co.kr/learn/courses/30/lessons/12903
// 가운데 글자 가져오기
// # 문제
// 단어 s의 가운데 글자를 반환하는 함수, solution을 만들어 보세요. 단어의 길이가 짝수라면 가운데 두글자를 반환하면 됩니다.

// # 제한
// s는 길이가 1 이상, 100이하인 스트링입니다.

package level1

import "fmt"

func q12903() {

	solution := func(s string) string {
		if (len(s) & 1) != 0 {
			return string(s[len(s)/2])
		}
		return string(s[len(s)/2-1]) + string(s[len(s)/2])
		// return s[int(float64(len(s)) / 2 - 0.5):int(float64(len(s)) / 2 + 1)]
	}

	fmt.Println(solution("abcde")) // "c"
	fmt.Println(solution("qwer"))  // "we"
}
