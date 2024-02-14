package code

import "fmt"

// q02739 is a function to solve the problem "구구단".
// source: https://www.acmicpc.net/problem/2739
func q02739() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= 9; i++ {
		fmt.Println(n, "*", i, "=", n*i)
	}
}
