package code

import (
	"bufio"
	"os"
)

// q10798 is a function to solve the problem "세로읽기".
// source: https://www.acmicpc.net/problem/10798
func q10798() {
	var words [5][]byte
	buf := make([]byte, 75) // 15 * 5
	reader := bufio.NewReader(os.Stdin)
	for i := range words {
		buf, _ = reader.ReadBytes('\n')
		words[i] = buf[:len(buf)-1]
	}
	result := make([]byte, 0, 75) // 15 * 5
	for i := 0; i < 15; i++ {
		for _, word := range words {
			if i < len(word) {
				result = append(result, word[i])
			}
		}
	}
	os.Stdout.Write(append(result, '\n'))
}
