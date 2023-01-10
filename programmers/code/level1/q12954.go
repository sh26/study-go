// https://school.programmers.co.kr/learn/courses/30/lessons/12954
// x만큼 간격이 있는 n개의 숫자
// # 문제
// 함수 solution은 정수 x와 자연수 n을 입력 받아,
// x부터 시작해 x씩 증가하는 숫자를 n개 지니는 리스트를 리턴해야 합니다.
// 다음 제한 조건을 보고, 조건을 만족하는 함수, solution을 완성해주세요.

// # 제한
// x는 -10000000 이상, 10000000 이하인 정수입니다.
// n은 1000 이하인 자연수입니다.

package level1

import "fmt"

func q12954() {

	solution := func(x int, n int) []int {
		i, res := 0, make([]int, n)
		for t := x; i < n; t, i, res[i] = t+x, i+1, t { }
		return res

		// ret[0] = x
		// for i := 1; i < n; i++ {
		//	ret[i] = ret[i-1] + ret[0]
		// }
	}

	fmt.Println(solution(2, 5))  // [2,4,6,8,10]
	fmt.Println(solution(4, 3))  // [4,8,12]
	fmt.Println(solution(-4, 2)) // [-4, -8]
}
