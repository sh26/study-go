// https://school.programmers.co.kr/learn/courses/30/lessons/12919
// 서울에서 김서방 찾기
// # 문제
// String형 배열 seoul의 element중 "Kim"의 위치 x를 찾아,
// "김서방은 x에 있다"는 String을 반환하는 함수, solution을 완성하세요.
// seoul에 "Kim"은 오직 한 번만 나타나며 잘못된 값이 입력되는 경우는 없습니다.

// # 제한
// seoul은 길이 1 이상, 1000 이하인 배열입니다.
// seoul의 원소는 길이 1 이상, 20 이하인 문자열입니다.
// "Kim"은 반드시 seoul 안에 포함되어 있습니다.

package level1

import (
	"fmt"
	"strings"
)

func q12919() {

	solution := func(seoul []string) string {
		str := strings.Join(seoul, " ")
		return fmt.Sprintf("김서방은 %d에 있다", strings.Count(str[:strings.Index(str, "Kim")], " "))
	}

	fmt.Println(solution([]string{"Jane", "Kim"}))
}
