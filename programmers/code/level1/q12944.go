// https://school.programmers.co.kr/learn/courses/30/lessons/12944
// 평균 구하기
// # 문제
// 정수를 담고 있는 배열 arr의 평균값을 return하는 함수, solution을 완성해보세요.

// # 제한
// arr은 길이 1 이상, 100 이하인 배열입니다.
// arr의 원소는 -10,000 이상 10,000 이하인 정수입니다.

package level1

import "fmt"

func q12944() {

	solution := func(arr []int) (avg float64) {
		for _, val := range arr {
			avg += float64(val)
		}
		avg /= float64(len(arr))
		return
	}

	fmt.Println(solution([]int{1, 2, 3, 4}))
	fmt.Println(solution([]int{5, 5}))
}
