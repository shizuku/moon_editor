package buffer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	st "github.com/shizuku/moon_editor/string"
)

type Buffer struct {
	data []st.String
}

func New() Buffer {
	return Buffer{
		data: make([]st.String, 0),
	}
}
func (b *Buffer) Read(fileName string) error {
	var s st.String
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
		s.Append(ch)
	}
	b.data = st.Split(s, '\n')
	return nil
}
func (b *Buffer) Write(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer f.Close()
	if err != nil {
		return err
	}
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
	var r st.String
	var l LineNumberer
	l.Init(b.LineNum())
	for i, v := range b.data {
		r.AppendString(v)
		if i != len(b.data)-1 {
			r.Append('\n')
		}
	}
	return r.String()
}
func (b *Buffer) Print() {
	fmt.Println("------------------------------------------------------")
	var r st.String
	l := NewLnr(b.LineNum())
	for i, v := range b.data {
		r.Append([]rune(l.Number(i+1))...)
		r.AppendString(v)
		r.Append('\n')
	}
	fmt.Print(r.String())
	fmt.Println("------------------------------------------------------")
}
func (b *Buffer) Insert(text *st.String, index int) {
	b.data = append(b.data[:index], append([]st.String{*text}, b.data[index:]...)...)
}
func (b *Buffer) Delete(index int) int {
	b.data = append(b.data[:index], b.data[index+1:]...)
	if len(b.data) == 0 {
		b.data = []st.String{}
	}
	if index - 1 <= 0 {
		return 0
	} else {
		return index - 1
	}
}
func (b *Buffer) Find(text *st.String, start int) (i int, j int) {
	for i, v := range b.data {
		if i < start {
			continue
		} else {
			j := v.Find(text)
			if j >= 0 {
				return i, j
			}
		}
	}
	return -1, -1
}
func (b *Buffer) Change(text, rp *st.String, lineIdx int) {
	b.data[lineIdx] = (b.data[lineIdx].Replace(text, rp))
}
