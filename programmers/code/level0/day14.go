package level0

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120890
// 가까운 수
func q120890() {

	solution := func(array []int, n int) int {
		array = append(array, n)         // 배열에 값 추가
		sort.Ints(array)                 // 정렬
		idx := sort.SearchInts(array, n) // value to index

		switch idx {
		case 0: // 가장 작음
			return array[idx+1]
		case len(array) - 1: // 가장 큼
			return array[idx-1]
		}

		if array[idx+1]-array[idx] <= array[idx]-array[idx-1] {
			// 다음 인덱스의 값이 더 가까움
			return array[idx+1]
		} else {
			// 이전 인덱스의 값이 더 가까움
			return array[idx-1]
		}
	}

	fmt.Println(solution([]int{3, 10, 28}, 22))  // 28
	fmt.Println(solution([]int{10, 11, 12}, 13)) // 12
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120891
// 369게임
func q120891() {

	solution := func(order int) int {
		return len(strings.NewReplacer("0", "", "1", "", "2", "", "4", "", "5", "", "7", "", "8", "").Replace(strconv.Itoa(order)))
	}

	fmt.Println(solution(3))
	fmt.Println(solution(29423))
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120892
// 암호 해독
func q120892() {

	solution := func(cipher string, code int) (word string) {

		for idx := code; idx <= len(cipher); idx += code {
			word += string(cipher[idx-1])
		}

		return
	}

	fmt.Println(solution("dfjardstddetckdaccccdegk", 4)) //	"attack"
	fmt.Println(solution("pfqallllabwaoclk", 2))         //	"fallback"
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120893
// 대문자와 소문자
func q120893() {

	solution := func(my_string string) string {
		str := make([]byte, len(my_string))
		for idx, chr := range my_string {
			if chr < 97 {
				str[idx] = byte(chr + 32)
			} else {
				str[idx] = byte(chr - 32)
			}
		}
		return string(str)
	}
	fmt.Println(solution("aA"))
	fmt.Println(solution("cccCCC"))     //"CCCccc"
	fmt.Println(solution("abCdEfghIJ")) //	"ABcDeFGHij"
}
