// https://www.acmicpc.net/problem/2588
// # 문제
// (세 자리 수) × (세 자리 수)는 다음과 같은 과정을 통하여 이루어진다. [https://www.acmicpc.net/upload/images/f5NhGHVLM4Ix74DtJrwfC97KepPl27s%20(1).png]
// (1)과 (2)위치에 들어갈 세 자리 자연수가 주어질 때 (3), (4), (5), (6)위치에 들어갈 값을 구하는 프로그램을 작성하시오.

// # 입력
// 첫째 줄에 (1)의 위치에 들어갈 세 자리 자연수가, 둘째 줄에 (2)의 위치에 들어갈 세자리 자연수가 주어진다.

// # 출력
// 첫째 줄부터 넷째 줄까지 차례대로 (3), (4), (5), (6)에 들어갈 값을 출력한다.

// # 예제
// [472,385] → [2360,3776,1416,181720]

package code

import (
	"fmt"
)

func q02558() {

	base, array, full := 0, "", []int{}

	fmt.Scanf("%d\n%s", &base, &array)

	for _, i := range array {
		full = append(full, int(i-'0'))
	}

	calc, digit := 0, 1

	for i := len(full) - 1; i >= 0; i-- {
		digit *= 10
		calc += full[i] * base * digit
		fmt.Println(full[i] * base)
	}

	fmt.Println(calc / 10)
}
