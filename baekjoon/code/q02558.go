package code

import (
	"fmt"
)

// q02558 is a function to solve the problem "곱셈".
// source: https://www.acmicpc.net/problem/2588
func q02558() {
	base, array, full := 0, "", []int{}

	fmt.Scanf("%d\n%s", &base, &array)

	for _, i := range array {
		full = append(full, int(i-'0'))
	}

	calc, digit := 0, 1

	for i := len(full) - 1; i >= 0; i-- {
		digit *= 10
		calc += full[i] * base * digit
		fmt.Println(full[i] * base)
	}

	fmt.Println(calc / 10)
}
