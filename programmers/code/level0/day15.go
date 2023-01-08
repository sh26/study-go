package level0

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120894
// 영어가 싫어요
func q120894() {

	solution := func(numbers string) int64 {
		num, _ := strconv.Atoi(strings.NewReplacer("zero", "0", "one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9").Replace(numbers))
		return int64(num)
	}

	fmt.Println(solution("onetwothreefourfivesixseveneightnine")) //	123456789
	fmt.Println(solution("onefourzerosixseven"))                  //	14067
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120895
// 인덱스 바꾸기
func q120895() {

	solution := func(my_string string, num1 int, num2 int) string {
		swap := []rune(my_string)
		swap[num1], swap[num2] = swap[num2], swap[num1]
		return string(swap)
	}

	fmt.Println(solution("hello", 1, 2))      // "hlelo"
	fmt.Println(solution("I love you", 3, 6)) // "I l veoyou"
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120896
// 한 번만 등장한 문자
func q120896() {

	solution := func(s string) string {
		once := []string{}
		for _, str := range s {
			if strings.Count(s, string(str)) == 1 {
				once = append(once, string(str))
			}
		}
		sort.Strings(once)
		return strings.Join(once, "")
	}

	fmt.Println(solution("abcabcadc")) // "d"
	fmt.Println(solution("abdc"))      // "abcd"
	fmt.Println(solution("hello"))     // "eho"
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120897
// 약수 구하기
func q120897() {

	solution := func(n int) (nums []int) {
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				nums = append(nums, i)
			}
		}
		return
	}

	fmt.Println(solution(24)) // [1, 2, 3, 4, 6, 8, 12, 24]
	fmt.Println(solution(29)) // [1, 29]
}
