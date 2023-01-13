package code

import (
	"baekjoon/code/step01"
	"baekjoon/code/step02"
	"baekjoon/code/step03"
	"baekjoon/code/step04"
	"baekjoon/code/step05"
	"baekjoon/code/step06"
)

func Run(step int) {
	switch step {
	case 1:
		step01.Answer()
	case 2:
		step02.Answer()
	case 3:
		step03.Answer()
	case 4:
		step04.Answer()
	case 5:
		step05.Answer()
	case 6:
		step06.Answer()
	}
}
