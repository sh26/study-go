package code

import (
	"io"
	"os"
)

// q11718 is a function to solve the problem "그대로 출력하기".
// source: https://www.acmicpc.net/problem/11718
func q11718() {
	io.Copy(os.Stdout, os.Stdin)
}
