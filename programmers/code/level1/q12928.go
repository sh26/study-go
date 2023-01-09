// https://school.programmers.co.kr/learn/courses/30/lessons/12928

package level1

import "fmt"

func q12928() {
	solution := func(n int) (sum int) {
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				sum += i
			}
		}
		return
	}

	fmt.Println(solution(12)) // 28
	fmt.Println(solution(5))  // 6
}
