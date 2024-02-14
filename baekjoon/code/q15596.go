package code

import "fmt"

// q15596 is a function to solve the problem "정수 N개의 합".
// source: https://www.acmicpc.net/problem/15596
func q15596() {
	sum := func(a []int) int {
		r := 0
		for _, n := range a {
			r += n
		}
		return r
	}

	fmt.Print(sum([]int{1}))
}
