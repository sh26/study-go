package code

import (
	"fmt"
)

// q01001 is a function to solve the problem "A-B".
// source: https://www.acmicpc.net/problem/1001
func q01001() {
	var i, j int
	fmt.Scanf("%d %d", &i, &j)
	fmt.Print(i - j)
}
