package level0

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120864
// 숨어있는 숫자의 덧셈 (2)
func q120864() {

	solution := func(my_string string) (sum int) {
		my_string = regexp.MustCompile("[a-zA-Z]").ReplaceAllString(my_string, " ")

		for _, v := range strings.Split(my_string, " ") {
			if v == "" {
				continue
			}
			n, _ := strconv.Atoi(string(v))
			sum += n
		}
		return
	}

	fmt.Println(solution("aAb1B2cC34oOp")) // 37
	fmt.Println(solution("1a2b3c4d123Z"))  // 133
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120866
// 안전지대
func q120866() {

	solution := func(board ...[]int) (safe int) {
		safe = len(board) * len(board)

		unsafe := func(xpos, ypos int) { // unsafe
			if board[xpos][ypos] != 1 {
				board[xpos][ypos] = -1
			}
		}

		for x, line := range board { // check
			for y, area := range line {
				if area == 0 || area == -1 {
					continue
				}

				if y > 0 { // 현재 라인
					unsafe(x, y-1)
				}
				if y < len(board)-1 {
					unsafe(x, y+1)
				}

				if x > 0 { // 위 라인
					unsafe(x-1, y)
					if y > 0 {
						unsafe(x-1, y-1)
					}
					if y < len(board)-1 {
						unsafe(x-1, y+1)
					}
				}

				if x < len(board)-1 { // 아래 라인
					unsafe(x+1, y)
					if y > 0 {
						unsafe(x+1, y-1)
					}
					if y < len(board)-1 {
						unsafe(x+1, y+1)
					}
				}
			}
		}

		for _, line := range board { // count
			for _, area := range line {
				if area == 0 {
					continue
				}
				safe--
			}
		}
		return
	}

	fmt.Println(solution([]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, []int{0, 0, 1, 0, 0}, []int{0, 0, 0, 0, 0}))                                         // 16
	fmt.Println(solution([]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, []int{0, 0, 1, 1, 0}, []int{0, 0, 0, 0, 0}))                                         // 13
	fmt.Println(solution([]int{1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1})) // 0
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120868
// 삼각형의 완성조건 (2)
func q120868() {

	solution := func(sides []int) int {
		sort.Ints(sides)
		return 2*sides[0] - 1
	}

	fmt.Println(solution([]int{1, 2}))  // 1
	fmt.Println(solution([]int{3, 6}))  // 5
	fmt.Println(solution([]int{11, 7})) // 13
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120869
// 외계어 사전
func q120869() {

	solution := func(spell []string, dic []string) int {

		check := func(word string) (pass bool) { // 검토
			pass = true
			for _, v := range spell {
				if !strings.Contains(word, v) {
					pass = false
				}
			}
			return
		}

		for _, word := range dic {
			if check(word) {
				return 1
			}
		}
		return 2
	}

	fmt.Println(solution([]string{"p", "o", "s"}, []string{"sod", "eocd", "qixm", "adio", "soo"}))       // 2
	fmt.Println(solution([]string{"z", "d", "x"}, []string{"def", "dww", "dzx", "loveaw"}))              // 1
	fmt.Println(solution([]string{"s", "o", "m", "d"}, []string{"moos", "dzx", "smm", "sunmmo", "som"})) // 2
}
