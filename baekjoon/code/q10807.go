package code

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// q10807 is a function to solve the problem "개수 세기".
// source: https://www.acmicpc.net/problem/10807
func q10807() {
	r, w := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	scanText := func() string {
		r.Scan()
		return r.Text()
	}
	r.Scan()
	fmt.Fprintln(w, strings.Count("?"+strings.ReplaceAll(scanText(), " ", " ?")+" ", "?"+scanText()+" "))
}
