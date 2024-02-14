package code

import (
	"fmt"
)

// q10998 is a function to solve the problem "A×B".
// source: https://www.acmicpc.net/problem/10998
func q10998() {
	var i, j int
	fmt.Scanf("%d %d", &i, &j)
	fmt.Print(i * j)
}
