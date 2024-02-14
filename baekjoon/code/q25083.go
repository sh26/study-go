package code

import (
	"bufio"
	"fmt"
	"os"
)

// q25083 is a function to solve the problem "새싹".
// source: https://www.acmicpc.net/problem/25083
func q25083() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fprintln(w, "         ,r'\"7")
	fmt.Fprintln(w, "r`-_   ,'  ,/")
	fmt.Fprintln(w, " \\. \". L_r'")
	fmt.Fprintln(w, "   `~\\/")
	fmt.Fprintln(w, "      |")
	fmt.Fprintln(w, "      |")
}
