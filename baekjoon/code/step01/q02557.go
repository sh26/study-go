// https://www.acmicpc.net/problem/2557
// # 문제
// Hello World!를 출력하시오.

// # 출력
// Hello World!를 출력하시오.

package step01

import (
	"bufio"
	"fmt"
	"os"
)

func q02557() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(w, "Hello World!")
}
