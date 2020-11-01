package str

type String struct {
	data []rune
}

func New() String {
	return String{}
}

func (s *String) Len() int {
	return len(s.data)
}

func (s *String) Clear() {
	s.data = []rune{}
}

func (s *String) Append(r rune) {
	s.data = append(s.data, r)
}

func (s *String) String() string {
	return string(s.data)
}

func (s *String) Split(sp rune) []String {
	var r []String
	var ss String
	for _, ch := range s.data {
		if ch == sp {
			r = append(r, ss)
			ss.Clear()
		} else {
			ss.Append(ch)
		}
	}
	r = append(r, ss)
	return r
}
