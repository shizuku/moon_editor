package string

type String struct{
	data []rune
}

func New(rs []rune) *String {
	return &String{data:rs}
}
func (s *String) Len() int {
	return len(s.data)
}
func (s *String) Append(r ...rune) {
	s.data = append(s.data, r...)
}
func (s *String) AppendString(r String) {
	s.data = append(s.data, r.data...)
}
func (s *String) Find(mod *String) int {
	var (
		a  int = 0
		da int = 0
		db int = 0
	)
	for a < s.Len() {
		if db >= mod.Len() {
			return a
		}
		if a+da >= s.Len() {
			return -1
		}
		if s.data[a+da] == mod.data[db] {
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
func (s *String) Replace(mod, rp *String) String {
	var (
		r  String
		a  int = 0
		da int = 0
		db int = 0
	)
	for a < s.Len() {
		if db >= mod.Len() {
			a += da
			da = 0
			db = 0
			r.data = append(r.data, rp.data...)
			continue
		}
		if a+da >= s.Len() {
			r.data = append(r.data, rp.data...)
			return r
		}
		if s.data[a+da] == mod.data[db] {
			da++
			db++
		} else {
			da = 0
			db = 0
			r.data = append(r.data, s.data[a])
			a++
		}
	}
	return r
}
func (s *String) String() string {
	return string(s.data)
}
func Split(str String, sp rune) []String {
	var r []String
	var s []rune
	for _, ch := range str.data {
		if ch == sp {
			r = append(r, String{data:s})
			s = []rune{}
		} else {
			s = append(s, ch)
		}
	}
	r = append(r, String{data:s})
	return r
}
