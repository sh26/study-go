// https://www.acmicpc.net/problem/25083
package code

// 아래 예제와 같이 새싹을 출력하시오.

/*
         ,r'"7
r`-_   ,'  ,/
 \. ". L_r'
   `~\/
      |
      |
*/

import (
	"bufio"
	"fmt"
	"os"
)

func q25083() {

	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(w, "         ,r'\"7")
	fmt.Fprintln(w, "r`-_   ,'  ,/")
	fmt.Fprintln(w, " \\. \". L_r'")
	fmt.Fprintln(w, "   `~\\/")
	fmt.Fprintln(w, "      |")
	fmt.Fprintln(w, "      |")
	w.Flush()
}
