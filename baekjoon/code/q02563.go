package code

import (
	"bufio"
	"fmt"
	"os"
)

// q02563 is a function to solve the problem "색종이".
// source: https://www.acmicpc.net/problem/2563
func q02563() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	type pos struct {
		x, y int
	}

	paper, size, p := map[pos]bool{}, 0, pos{}
	fmt.Scanf("%d\n", &size)
	for idx := 0; idx < size; idx++ {
		fmt.Scanf("%d %d\n", &p.x, &p.y)
		for i := p.x + 10; p.x < i; p.x++ {
			for j := p.y + 10; p.y < j; p.y++ {
				paper[p] = true
			}
			p.y -= 10
		}
	}
	size = 0
	for p.x = 0; p.x < 100; p.x++ {
		for p.y = 0; p.y < 100; p.y++ {
			if paper[p] {
				size++
			}
		}
	}

	fmt.Fprintln(w, size)
}
