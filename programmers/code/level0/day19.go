package level0

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120583
// 중복된 숫자 개수
func q120583() {

	solution := func(array []int, n int) int {
		return strings.Count("!"+strings.NewReplacer("[", "", "]", "", " ", "?!").Replace(fmt.Sprint(array)+"?"), "!"+fmt.Sprint(n)+"?")
	}

	fmt.Println(solution([]int{100, 100, 100, 100, 100, 100}, 100)) // 2
	fmt.Println(solution([]int{0, 2, 3, 4}, 2))                     // 0
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120585
// 머쓱이보다 키 큰 사람
func q120585() {

	solution := func(array []int, height int) int {
		array = append(array, height)
		sort.Ints(array)
		return int(math.Max(0, float64(len(array)-sort.SearchInts(array, height+1))))
	}

	fmt.Println(solution([]int{149, 180, 192, 170}, 167))              // 3
	fmt.Println(solution([]int{180, 120, 140}, 190))                   // 0
	fmt.Println(solution([]int{1, 1, 1, 1, 1, 100, 100, 100, 100}, 1)) // 4
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120912
// 7의 개수
func q120912() {

	solution := func(array []int) int {
		return strings.Count(fmt.Sprint(array), "7")
	}

	fmt.Println(solution([]int{7, 77, 17})) // 4
	fmt.Println(solution([]int{10, 29}))    // 0
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120913
// 잘라서 배열로 저장하기
func q120913() {

	solution := func(my_str string, n int) (res []string) {
		for i := 0; i < len(my_str); i += n {
			res = append(res, my_str[i:int(math.Min(float64(i+n), float64(len(my_str))))])
		}
		return
	}

	fmt.Println(solution("abc1Addfggg4556b", 6)) // ["abc1Ad", "dfggg4", "556b"]
	fmt.Println(solution("abcdef123", 3))        // ["abc", "def", "123"]
}
