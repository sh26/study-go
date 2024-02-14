package code

import (
	"fmt"
)

// q10869 is a function to solve the problem "사칙연산".
// source: https://www.acmicpc.net/problem/10869
func q10869() {
	var i, j int
	fmt.Scanf("%d %d", &i, &j)
	fmt.Println(i + j)
	fmt.Println(i - j)
	fmt.Println(i * j)
	fmt.Println(i / j)
	fmt.Println(i % j)
}
