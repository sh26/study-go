// https://school.programmers.co.kr/learn/courses/30/lessons/12933
// 정수 내림차순으로 배치하기
// # 문제
// 함수 solution은 정수 n을 매개변수로 입력받습니다.
// n의 각 자릿수를 큰것부터 작은 순으로 정렬한 새로운 정수를 리턴해주세요.
// 예를들어 n이 118372면 873211을 리턴하면 됩니다.

// # 제한
// n은 1이상 8000000000 이하인 자연수입니다.

package level1

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func q12933() {

	solution := func(n int64) int64 {
		slice, res := strings.Split(fmt.Sprint(n), ""), int64(0)
		sort.Strings(slice)
		for idx := range slice {
			tmp, _ := strconv.ParseInt(slice[idx], 10, 64)
			res += int64(math.Pow10(idx)) * tmp
		}
		return res
	}

	fmt.Println(solution(118372)) // 873211
}
