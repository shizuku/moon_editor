package buffer

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Buffer struct {
	data []string
}

func New() Buffer {
	return Buffer{
		data: []string{},
	}
}
func (b *Buffer) Read(fileName string) error {
	var s []rune
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		return err
	}
	rd := bufio.NewReader(f)
	for {
		ch, _, er := rd.ReadRune()
		if er != nil {
			if er == io.EOF {
				break
			} else {
				return er
			}
		}
		s = append(s, ch)
	}
	b.data = split(string(s), '\n')
	return nil
}
func (b *Buffer) Write(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer f.Close()
	if err != nil {
		return err
	}
	//_, err = f.Write([]byte(b.String()))
	_, err = f.WriteString(b.String())
	if err != nil {
		return err
	}
	return nil
}
func (b *Buffer) LineNum() int {
	return len(b.data)
}
func (b *Buffer) String() string {
	var r []rune
	var l LineNumberer
	l.Init(b.LineNum())
	for i, v := range b.data {
		r = append(r, []rune(v)...)
		if i != len(b.data)-1 {
			r = append(r, '\n')
		}
	}
	return string(r)
}
func (b *Buffer) Print() {
	fmt.Println("------------------------------------------------------")
	var r []rune
	var l LineNumberer
	l.Init(b.LineNum())
	for i, v := range b.data {
		r = append(r, []rune(l.Number(i+1))...)
		r = append(r, []rune(v)...)
		r = append(r, '\n')
	}
	fmt.Print(string(r))
	fmt.Println("------------------------------------------------------")
}
func (b *Buffer) Insert(text string, index int) {
	b.data = append(b.data[:index], append([]string{text}, b.data[index:]...)...)
}
func (b *Buffer) Delete(index int) {
	b.data = append(b.data[:index], b.data[index+1:]...)
}
func (b *Buffer) Find(text string) (i int, j int) {
	for i, v := range b.data {
		j := strFind(ru(v), ru(text))
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}
func (b *Buffer) Change(text, rp string, lineIdx int) {
	b.data[lineIdx] = st(strReplace(ru(b.data[lineIdx]), ru(text), ru(rp)))
}

func ru(s string) []rune {
	return []rune(s)
}
func st(r []rune) string {
	fmt.Println(r)
	return string(r)
}
func split(str string, sp rune) []string {
	var r []string
	var s []rune
	for _, ch := range str {
		if ch == sp {
			r = append(r, string(s))
			s = []rune{}
		} else {
			s = append(s, ch)
		}
	}
	r = append(r, string(s))
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
			r = append(r, []rune(rp)...)
			continue
		}
		if a+da >= len(s) {
			r = append(r, []rune(rp)...)
			return r
		}
		if s[a+da] == mod[db] {
			da++
			db++
		} else {
			da = 0
			db = 0
			r = append(r, rune(s[a]))
			a++
		}
	}
	return r
}
