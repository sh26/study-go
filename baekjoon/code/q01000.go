package code

import (
	"fmt"
)

// q01000 is a function to solve the problem "A+B".
// source: https://www.acmicpc.net/problem/1000
func q01000() {
	var i, j int
	fmt.Scanf("%d %d", &i, &j)
	fmt.Print(i + j)
}
