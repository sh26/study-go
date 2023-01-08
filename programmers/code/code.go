package code

import (
	"programmers/code/level0"
	"programmers/code/level1"
)

func Run(level int) {
	switch level {
	case 0:
		level0.Answer()
	case 1:
		level1.Answer()
	}
}
