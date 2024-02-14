package code

import (
	"fmt"
)

// q01008 is a function to solve the problem "A/B".
// source: https://www.acmicpc.net/problem/1008
func q01008() {
	var i, j float64
	fmt.Scanf("%f %f", &i, &j)
	fmt.Print(i / j)
}
