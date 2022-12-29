// https://www.acmicpc.net/problem/2738
// # 문제
// N*M크기의 두 행렬 A와 B가 주어졌을 때, 두 행렬을 더하는 프로그램을 작성하시오.

// # 입력
// 첫째 줄에 행렬의 크기 N 과 M이 주어진다.
// 둘째 줄부터 N개의 줄에 행렬 A의 원소 M개가 차례대로 주어진다.
// 이어서 N개의 줄에 행렬 B의 원소 M개가 차례대로 주어진다.
// N과 M은 100보다 작거나 같고, 행렬의 원소는 절댓값이 100보다 작거나 같은 정수이다.

// # 출력
// 첫째 줄부터 N개의 줄에 행렬 A와 B를 더한 행렬을 출력한다. 행렬의 각 원소는 공백으로 구분한다.

// # 예제
// [3 3],[1 1 1],[2 2 2],[0 1 0],[3 3 3],[4 4 4],[5 5 100] → [4 4 4],[6 6 6],[5 6 100]

package step02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func q02738() {
	r, w, size := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout), []int{0, 0}
	defer w.Flush()

	fmt.Fscanf(r, "%d %d\n", &size[0], &size[1])

	procs := make([][]int, size[0])

	for i := 0; i < size[0]; i++ {
		procs[i] = make([]int, size[1])
		line, _ := r.ReadString('\n')

		for j, str := range strings.Split(strings.TrimSpace(line), " ") {
			procs[i][j], _ = strconv.Atoi(str)
		}
	}

	for i := 0; i < size[0]; i++ {
		line, _ := r.ReadString('\n')
		for j, str := range strings.Split(strings.TrimSpace(line), " ") {
			num, _ := strconv.Atoi(str)
			procs[i][j] += num
		}
		str := fmt.Sprint(procs[i])

		fmt.Fprintln(w, str[1:len(str)-1])
	}
}
