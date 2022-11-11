package code

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120854
// 배열 원소의 길이
func q120854() {

	solution := func(strlist []string) (strlen []int) {
		for _, v := range strlist {
			strlen = append(strlen, len(v))
		}

		return
	}

	fmt.Println(solution([]string{"We", "are", "the", "world!"})) // [2, 3, 3, 6]
	fmt.Println(solution([]string{"I", "Love", "Programmers."}))  // [1, 4, 12]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120860
// 직사각형 넓이 구하기
func q120860() {

	solution := func(dots ...[]int) int {
		xmin, xmax, ymin, ymax := 256, -256, 256, -256

		for _, dot := range dots {
			if dot[0] < xmin {
				xmin = dot[0]
			}
			if dot[0] > xmax {
				xmax = dot[0]
			}
			if dot[1] < ymin {
				ymin = dot[1]
			}
			if dot[1] > ymax {
				ymax = dot[1]
			}
		}

		return (xmax - xmin) * (ymax - ymin)
	}

	fmt.Println(solution([]int{1, 1}, []int{2, 1}, []int{2, 2}, []int{1, 2}))     // 1
	fmt.Println(solution([]int{-1, -1}, []int{1, 1}, []int{1, -1}, []int{-1, 1})) // 4
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120861
// 캐릭터의 좌표
func q120861() {

	solution := func(keyinput []string, board []int) []int {

		position, width, height := []int{0, 0}, board[0]/2, board[1]/2

		for _, v := range keyinput {
			if position[0] > -width && v == "left" {
				position[0]--
			} else if position[0] < width && v == "right" {
				position[0]++
			}
			if position[1] < height && v == "up" {
				position[1]++
			} else if position[1] > -height && v == "down" {
				position[1]--
			}
		}

		return position
	}

	fmt.Println(solution([]string{"left", "right", "up", "right", "right"}, []int{11, 11})) // [2, 1]
	fmt.Println(solution([]string{"down", "down", "down", "down", "down"}, []int{7, 9}))    // [0, -4]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120862
// 최댓값 만들기 (2)
func q120862() {

	solution := func(numbers []int) int {
		sort.Ints(numbers)
		if numbers[len(numbers)-1]*numbers[len(numbers)-2] > numbers[0]*numbers[1] {
			return numbers[len(numbers)-1] * numbers[len(numbers)-2]
		}
		return numbers[0] * numbers[1]
	}

	fmt.Println(solution([]int{1, 2, -3, 4, -5}))         // 15
	fmt.Println(solution([]int{0, -31, 24, 10, 1, 9}))    // 240
	fmt.Println(solution([]int{10, 20, 30, 5, 5, 20, 5})) // 600
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120863
// 다항식 더하기
func q120863() {

	solution := func(polynomial string) string {

		x, n := 0, 0

		for _, v := range strings.Split(polynomial, " + ") {
			if strings.Contains(v, "x") {
				v = strings.Replace(v, "x", "", -1)
				if v != "" {
					c, _ := strconv.Atoi(v)
					x += c
				} else {
					x++
				}
			} else {
				c, _ := strconv.Atoi(v)
				n += c
			}
		}

		if x == 0 {
			if n == 0 {
				return "0"
			}
			return fmt.Sprint(n)
		}

		if n == 0 {
			if x == 1 {
				return "x"
			}
			return fmt.Sprintf("%dx", x)
		}

		if x == 1 {
			return fmt.Sprintf("x + %d", n)
		}

		return fmt.Sprintf("%dx + %d", x, n)
	}

	fmt.Println(solution("3x + 7 + x")) // "x"
	fmt.Println(solution("x + x + x"))  // "3x"
}

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
	fmt.Println(solution("90a9aaA"))       // 99
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120866
// 안전지대
func q120866() {

	solution := func(board [][]int) (safe int) {
		safe = len(board) * len(board)

		// unsafe
		unsafe := func(xpos, ypos int) {
			if board[xpos][ypos] != 1 {
				board[xpos][ypos] = -1
			}
		}

		// check
		for x, line := range board {
			for y, area := range line {
				if area == 0 || area == -1 {
					continue
				}

				// 현재 라인
				if y > 0 {
					unsafe(x, y-1)
				}
				if y < len(board)-1 {
					unsafe(x, y+1)
				}

				// 위 라인
				if x > 0 {
					unsafe(x-1, y)
					if y > 0 {
						unsafe(x-1, y-1)
					}
					if y < len(board)-1 {
						unsafe(x-1, y+1)
					}
				}

				// 아래 라인
				if x < len(board)-1 {
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

		// count
		for _, line := range board {
			for _, area := range line {
				if area == 0 {
					continue
				}
				safe--
			}
		}

		return
	}

	fmt.Println(solution([][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}}))                                    // 16
	fmt.Println(solution([][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 1, 1, 0}, {0, 0, 0, 0, 0}}))                                    // 13
	fmt.Println(solution([][]int{{1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}})) // 0
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120868
// 삼각형의 완성조건 (2)
func q120868() {

	solution := func(sides []int) int {
		sort.Ints(sides)
		return 2*sides[0] - 1
	}

	fmt.Println(solution([]int{1, 2}))  //	1
	fmt.Println(solution([]int{3, 6}))  //	5
	fmt.Println(solution([]int{11, 7})) // 13
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120869
// 외계어 사전
func q120869() {

	solution := func(spell []string, dic []string) int {

		// 검토
		check := func(word string) (pass bool) {
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

	fmt.Println(solution([]string{"p", "o", "s"}, []string{"sod", "eocd", "qixm", "adio", "soo"}))       //	2
	fmt.Println(solution([]string{"z", "d", "x"}, []string{"def", "dww", "dzx", "loveaw"}))              //	1
	fmt.Println(solution([]string{"s", "o", "m", "d"}, []string{"moos", "dzx", "smm", "sunmmo", "som"})) //	2
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120871
// 저주의 숫자 3
func q120871() {

	solution := func(n int) int {
		for dec, cnt := 1, 1; ; dec++ {
			// 3의 배수거나 10 이상이며 수에 3이 포함되는 경우
			for {
				if cnt%3 == 0 || strings.Contains(strconv.Itoa(cnt), "3") {
					cnt++
				} else {
					if dec >= n {
						return cnt
					}
					cnt++
					break
				}
			}
		}
	}

	fmt.Println(solution(15)) // 25
	fmt.Println(solution(40)) // 76
}
