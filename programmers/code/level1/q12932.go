// https://school.programmers.co.kr/learn/courses/30/lessons/12932
// 자연수 뒤집어 배열로 만들기
// # 문제
// 자연수 n을 뒤집어 각 자리 숫자를 원소로 가지는 배열 형태로 리턴해주세요. 예를들어 n이 12345이면 [5,4,3,2,1]을 리턴합니다.

// # 제한
// n은 10,000,000,000이하인 자연수입니다.

package level1

import "fmt"

func q12932() {

	solution := func(n int64) (num []int) {
		for n > 0 {
			num = append(num, int(n%10))
			n /= 10
		}
		return
	}

	fmt.Println(solution(12345))
}
