package buffer

func split(str []rune, sp rune) [][]rune {
	var r [][]rune
	var s []rune
	for _, ch := range str {
		if ch == sp {
			r = append(r, (s))
			s = []rune{}
		} else {
			s = append(s, ch)
		}
	}
	r = append(r, (s))
	return r
}
func strFind(s, mod []rune) int {
	var (
		a  int = 0
		da int = 0
		db int = 0
	)
	for a < len(s) {
		if db >= len(mod) {
			return a
		}
		if a+da >= len(s) {
			return -1
		}
		if s[a+da] == mod[db] {
			da++
			db++
		} else {
			da = 0
			db = 0
			a++
		}
	}
	return -1
}
func strReplace(s, mod, rp []rune) []rune {
	var (
		r  []rune
		a  int = 0
		da int = 0
		db int = 0
	)
	for a < len(s) {
		if db >= len(mod) {
			a += da
			da = 0
			db = 0
			r = append(r, rp...)
			continue
		}
		if a+da >= len(s) {
			r = append(r, rp...)
			return r
		}
		if s[a+da] == mod[db] {
			da++
			db++
		} else {
			da = 0
			db = 0
			r = append(r, s[a])
			a++
		}
	}
	return r
}
