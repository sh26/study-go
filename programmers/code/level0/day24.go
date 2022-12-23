package code

import (
	"fmt"
	"strconv"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120884
// 치킨 쿠폰
func q120884() {

	solution := func(chicken int) (answer int) {
		for chicken >= 10 {
			answer += chicken / 10
			chicken = (chicken / 10) + (chicken % 10)
		}

		return
	}

	fmt.Println(solution(100))  // 11
	fmt.Println(solution(1081)) // 120
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120885
// 이진수 더하기
func q120885() {

	solution := func(bin1 string, bin2 string) string {

		bin1num, bin2num := func() int64 { tmp, _ := strconv.ParseInt(bin1, 2, 64); return tmp }(), func() int64 { tmp, _ := strconv.ParseInt(bin2, 2, 64); return tmp }()

		return strconv.FormatInt(bin1num+bin2num, 2)
	}

	fmt.Println(solution("10", "11"))     // "101"
	fmt.Println(solution("1001", "1111")) // "11000"
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120886
// A로 B 만들기
func q120886() {

	solution := func(before string, after string) int {
		diff := map[rune]int{}
		for _, val := range before {
			diff[val]++
		}

		for _, val := range after {
			if strings.Count(after, string(val)) != diff[val] {
				return 0
			}
		}
		return 1
	}

	fmt.Println(solution("olleh", "hello")) // 1
	fmt.Println(solution("allpe", "apple")) // 0
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120887
// k의 개수
func q120887() {

	solution := func(i int, j int, k int) (cnt int) {
		for ; i <= j; i++ {
			temp := i
			for temp > 0 {
				modnum := temp % 10
				if modnum == k {
					cnt++
				}
				temp /= 10
			}
		}
		return
	}
	fmt.Println(solution(1, 13, 1))  // 6
	fmt.Println(solution(10, 50, 5)) // 5
	fmt.Println(solution(3, 10, 2))  // 0
}
