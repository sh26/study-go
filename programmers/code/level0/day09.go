package code

import (
	"bytes"
	"fmt"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120837
// 개미 군단
func q120837() {

	solution := func(hp int) int {
		return (hp / 5) + (hp%5)/3 + (hp%5%3)/1
	}

	fmt.Println(solution(23))  // 5
	fmt.Println(solution(24))  // 6
	fmt.Println(solution(999)) // 201
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120838
// 모스부호 (1)
func q120838() {

	solution := func(letter string) string {

		morse := map[string]string{".-": "a", "-...": "b", "-.-.": "c", "-..": "d", ".": "e", "..-.": "f",
			"--.": "g", "....": "h", "..": "i", ".---": "j", "-.-": "k", ".-..": "l",
			"--": "m", "-.": "n", "---": "o", ".--.": "p", "--.-": "q", ".-.": "r",
			"...": "s", "-": "t", "..-": "u", "...-": "v", ".--": "w", "-..-": "x",
			"-.--": "y", "--..": "z",
		}

		decode := bytes.Buffer{}

		for _, v := range strings.Split(letter, " ") {
			decode.WriteByte(morse[v][0])
		}

		return decode.String()
	}

	fmt.Println(solution(".... . .-.. .-.. ---"))    //"hello"
	fmt.Println(solution(".--. -.-- - .... --- -.")) //"python"
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120839
// 가위 바위 보
func q120839() {

	solution := func(rsp string) string {
		win := bytes.Buffer{}
		for _, v := range rsp {
			switch v {
			case '0':
				win.WriteRune('5')
			case '2':
				win.WriteRune('0')
			case '5':
				win.WriteRune('2')
			}
		}
		return win.String()
	}

	fmt.Println(solution("2"))   // "0"
	fmt.Println(solution("205")) //	"052"
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120840
// 구슬을 나누는 경우의 수
func q120840() {

	solution := func(balls int, share int) int {

		fact := func(n int) float64 {
			var v float64 = 1
			for i := 1; i <= n; i++ {
				v *= float64(i)
			}
			return v
		}

		// uint     either 32 or 64 bits
		// int      same size as uint
		return int(fact(balls) / (fact(balls-share) * fact(share)))
	}

	fmt.Println(solution(3, 30)) // 3
	fmt.Println(solution(5, 3))  // 10
}
