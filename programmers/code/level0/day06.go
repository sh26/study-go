package level0

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120822
// 문자열 뒤집기
func q120822() {

	solution := func(my_string string) string {

		runes := []rune(my_string)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		return string(runes)
	}

	fmt.Println(solution("jaron")) // "noraj"
	fmt.Println(solution("bread")) // "daerb"
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120823
// 직각삼각형 출력하기
func q120823() {

	solution := func() {
		var n int
		r := bufio.NewReader(os.Stdin)
		w := bufio.NewWriter(os.Stdout)

		fmt.Fscan(r, &n)

		for i := 0; i < n; i++ {
			for j := 0; j < i+1; j++ {
				fmt.Fprint(w, "*")
			}
			fmt.Fprintln(w)
		}
		w.Flush()
	}

	solution()
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120824
// 짝수 홀수 개수
func q120824() {

	solution := func(num_list ...int) []int {
		odd := 0
		for _, v := range num_list {
			if v&1 != 0 {
				odd++
			}
		}
		return []int{len(num_list) - odd, odd}
	}

	fmt.Println(solution(1, 2, 3, 4, 5)) //	[2, 3]
	fmt.Println(solution(1, 3, 5, 7))    // [0, 4]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120825
// 문자 반복 출력하기
func q120825() {

	solution := func(my_string string, n int) string {
		b := bytes.Buffer{}
		for _, v := range my_string {
			for i := n; i != 0; i-- {
				b.WriteString(string(v))
			}
		}

		return b.String()
	}

	fmt.Println(solution("hello", 3)) // "hhheeellllllooo"
}
