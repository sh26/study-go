package level0

import (
	"bytes"
	"fmt"
	"math"
	"sort"
)

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
// 진료 순서 정하기
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

// https://school.programmers.co.kr/learn/courses/30/lessons/120836
// 순서쌍의 개수
func q120836() {

	solution := func(n int) (cnt int) {

		sqrt := math.Sqrt(float64(n))
		sqrt_round := math.Round(sqrt)

		if sqrt == sqrt_round {
			if n%int(sqrt) == 0 {
				cnt++
			}
		} else {
			sqrt += 1
		}

		for i := 1; i < int(sqrt); i++ {
			if n%i == 0 {
				cnt += 2
			}
		}

		/*
			for i := 1; i*i <= n; i++ {
				if n%i == 0 {
					cnt += 2
					if i == n/i {
						cnt--
					}
				}
			}
		*/

		return cnt
	}

	fmt.Println(solution(20))  // 6
	fmt.Println(solution(100)) // 9answer
}
