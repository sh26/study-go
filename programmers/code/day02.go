package code

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120815
// 피자 나눠 먹기 (2)
func q120815() {

	solution := func(n int) int {
		return n / func(a, b int) int {
			for a > 0 {
				a, b = b%a, a
			}
			return b
		}(n, 6)
	}

	fmt.Println(solution(6))  // 1
	fmt.Println(solution(10)) // 5
	fmt.Println(solution(4))  // 2
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120816
// 피자 나눠 먹기 (3)
func q120816() {

	solution := func(slice int, n int) int {
		if n%slice != 0 {
			return n/slice + 1
		}
		return n / slice
		//    (n + slice - 1) / slice
	}

	fmt.Println(solution(7, 10)) // 2
	fmt.Println(solution(4, 12)) // 3
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120817
// 배열의 평균값
func q120817() {

	solution := func(numbers ...int) float64 {

		total := 0

		for _, v := range numbers {
			total += v
		}

		return float64(total) / float64(len(numbers))
	}

	fmt.Println(solution(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))              //	5.5
	fmt.Println(solution(89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99)) //	94.0
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120818
// 옷가게 할인 받기
func q120818() {

	solution := func(price int) int {

		if price >= 500000 {
			price = price * 80
		} else if price >= 300000 {
			price = price * 90
		} else if price >= 100000 {
			price = price * 95
		} else {
			price *= 100
		}

		return price / 100
	}

	fmt.Println(solution(150000)) //	142,500
	fmt.Println(solution(580000)) //	464,000
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120819
// 옷가게 할인 받기
func q120819() {

	solution := func(money int) []int {
		return []int{money / 5500, money % 5500}
	}

	fmt.Println(solution(5500))  // [1, 0]
	fmt.Println(solution(15000)) // [2, 4000]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120820
// 옷가게 할인 받기
func q120820() {

	solution := func(age int) int {
		return 2023 - age
	}

	fmt.Println(solution(40)) // 1983
	fmt.Println(solution(23)) // 2000
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120821
// 배열 뒤집기
func q120821() {

	solution := func(num_list ...int) []int {

		for i, j := 0, len(num_list)-1; i < j; i, j = i+1, j-1 {
			num_list[i], num_list[j] = num_list[j], num_list[i]
		}

		return num_list
	}

	fmt.Println(solution(1, 2, 3, 4, 5))       // [5, 4, 3, 2, 1]
	fmt.Println(solution(1, 1, 1, 1, 1, 2))    // [2, 1, 1, 1, 1, 1]
	fmt.Println(solution(1, 0, 1, 1, 1, 3, 5)) // [5, 3, 1, 1, 1, 0, 1]
}

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

// https://school.programmers.co.kr/learn/courses/30/lessons/120826
// 특정 문자 제거하기
func q120826() {

	solution := func(my_string string, letter string) string {
		return strings.Replace(my_string, letter, "", -1)
	}

	fmt.Println(solution("abcdef", "f")) //	"abcde"
	fmt.Println(solution("BCBdbe", "B")) //	"Cdbe
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120829
// 각도기
func q120829() {

	solution := func(angle int) int {

		if 0 < angle && angle < 90 {
			return 1
		} else if angle == 90 {
			return 2
		} else if 90 < angle && angle < 180 {
			return 3
		} else {
			return 4
		}
	}

	fmt.Println(solution(70))  // 1
	fmt.Println(solution(91))  // 3
	fmt.Println(solution(180)) // 4
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120830
// 양꼬치
func q120830() {

	solution := func(n int, k int) int {
		return 12000*n + 2000*(k-(n/10))
	}

	fmt.Println(solution(10, 3)) //	124,000
	fmt.Println(solution(64, 6)) //	768,000
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120831
// 짝수의 합
func q120831() {

	solution := func(n int) (total int) {
		for i := 0; i <= n; i += 2 {
			total += i
		}
		return
	}

	fmt.Println(solution(10)) // 30
	fmt.Println(solution(4))  // 6
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120833
// 배열 자르기
func q120833() {

	solution := func(numbers []int, num1 int, num2 int) []int {
		return numbers[num1 : num2+1]
	}

	fmt.Println(solution([]int{1, 2, 3, 4, 5}, 1, 3)) //	[2, 3, 4]
	fmt.Println(solution([]int{1, 3, 5}, 1, 2))       //	[3, 5]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120834
// 외계행성의 나이
func q120834() {

	solution := func(age int) string {
		str := "abcdefghij"

		b := bytes.Buffer{}

		for age != 0 {
			b.WriteByte(str[age%10])
			age /= 10
		}

		runes := []rune(b.String())
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		return string(runes)
	}

	fmt.Println(solution(23))  // "cd"
	fmt.Println(solution(51))  // "fb"
	fmt.Println(solution(100)) // "baa"
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120835
// 외계행성의 나이
func q120835() {

	solution := func(emergency ...int) []int {
		n := map[int]int{}

		for _, v := range emergency {
			n[v] = 0
		}

		rank := make([]int, len(emergency))
		copy(rank, emergency)

		sort.Sort(sort.Reverse(sort.IntSlice(rank)))

		for i, v := range rank {
			n[v] = i + 1
		}

		for i, v := range emergency {
			emergency[i] = n[v]
		}

		return emergency
	}

	fmt.Println(solution(3, 76, 24))           // [3, 1, 2]
	fmt.Println(solution(1, 2, 3, 4, 5, 6, 7)) // [7, 6, 5, 4, 3, 2, 1]
	fmt.Println(solution(30, 10, 23, 6, 100))  // [2, 4, 3, 5, 1]
}
