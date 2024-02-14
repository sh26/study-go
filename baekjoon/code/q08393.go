package code

import "fmt"

// q08393 is a function to solve the problem "합".
// source: https://www.acmicpc.net/problem/8393
func q08393() {
	var n int
	fmt.Scan(&n)
	fmt.Print(n * (n + 1) / 2)
}
