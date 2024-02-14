package code

import "fmt"

// q10430 is a function to solve the problem "나머지".
// source: https://www.acmicpc.net/problem/10430
func q10430() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Println((a + b) % c)
	fmt.Println(((a % c) + (b % c)) % c)
	fmt.Println((a * b) % c)
	fmt.Println(((a % c) * (b % c)) % c)
}
