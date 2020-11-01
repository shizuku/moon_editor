package buffer

import (
	"strconv"
)

type LineNumberer struct {
	p int
}

func (l *LineNumberer) Init(maxIdx int) {
	l.p = 1
	for {
		if maxIdx < 10 {
			break
		} else {
			l.p++
			maxIdx /= 10
		}
	}
}

func (l *LineNumberer) Number(idx int) string {
	var r []byte
	s := strconv.Itoa(idx)
	for i := 0; i < l.p-len(s); i++ {
		r = append(r, ' ')
	}
	r = append(r, s...)
	r = append(r, ':')
	return string(r)
}
