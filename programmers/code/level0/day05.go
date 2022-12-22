package code

import "fmt"

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
// 아이스 아메리카노
func q120819() {

	solution := func(money int) []int {
		return []int{money / 5500, money % 5500}
	}

	fmt.Println(solution(5500))  // [1, 0]
	fmt.Println(solution(15000)) // [2, 4000]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120820
// 나이 출력
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
