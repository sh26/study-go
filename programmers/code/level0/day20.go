package level0

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

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

	fmt.Println(solution([]string{"left", "right", "up", "right", "right"}, []int{11, 11})) // []int{2, 1}
	fmt.Println(solution([]string{"down", "down", "down", "down", "down"}, []int{7, 9}))    // []int{0, -4}
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

	fmt.Println(solution("3x + 7 + x")) // "4x + 7"
	fmt.Println(solution("x + x + x"))  // "3x"
}
