// https://www.acmicpc.net/problem/1065
// # 문제
// 어떤 양의 정수 X의 각 자리가 등차수열을 이룬다면, 그 수를 한수라고 한다. 
// 등차수열은 연속된 두 개의 수의 차이가 일정한 수열을 말한다. 
// N이 주어졌을 때, 1보다 크거나 같고, N보다 작거나 같은 한수의 개수를 출력하는 프로그램을 작성하시오. 

// # 입력
// 첫째 줄에 1,000보다 작거나 같은 자연수 N이 주어진다.

// # 출력
// 첫째 줄에 1보다 크거나 같고, N보다 작거나 같은 한수의 개수를 출력한다.

// # 예제
// [110] → [99], [1] → [1], [210] → [105], [1000] → [144], [500] → [119]

package step05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func q01065() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	defer w.Flush()

	scanInt := func() (num int) {
		r.Scan()
		num, _ = strconv.Atoi(r.Text())
		return
	}

	if limit := scanInt(); limit < 100 {
		fmt.Fprintln(w, limit)
	} else {
		cnt := 99
		for ; limit >= 100; limit-- {
			n1, n2, n3 := limit%10, limit/10%10, limit/100
			if n3-n2 == n2-n1 {
				cnt++
			}
		}
		fmt.Fprintln(w, cnt)
	}
}
