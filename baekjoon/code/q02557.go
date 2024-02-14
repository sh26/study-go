package code

import (
	"bufio"
	"fmt"
	"os"
)

// q02557 is a function to solve the problem "Hello World".
// source: https://www.acmicpc.net/problem/2557
func q02557() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(w, "Hello World!")
}
