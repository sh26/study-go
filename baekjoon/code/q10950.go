package code

import "fmt"

// q10950 is a function to solve the problem "A+B - 3".
// source: https://www.acmicpc.net/problem/10950
func q10950() {
	var n, a, b int
	fmt.Scan(&n)
	s := ""
	for 0 < n {
		n--
		fmt.Scan(&a, &b)
		s += fmt.Sprintln(a + b)
	}
	fmt.Print(s)
}
