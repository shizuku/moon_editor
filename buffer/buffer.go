package buffer

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Buffer struct {
	data [][]rune
}

func New() Buffer {
	return Buffer{
		data: [][]rune{},
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
	b.data = split(s, '\n')
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
	l := NewLnr(b.LineNum())
	for i, v := range b.data {
		r = append(r, []rune(l.Number(i+1))...)
		r = append(r, []rune(v)...)
		r = append(r, '\n')
	}
	fmt.Print(string(r))
	fmt.Println("------------------------------------------------------")
}
func (b *Buffer) Insert(text []rune, index int) {
	b.data = append(b.data[:index], append([][]rune{text}, b.data[index:]...)...)
}
func (b *Buffer) Delete(index int) {
	b.data = append(b.data[:index], b.data[index+1:]...)
}
func (b *Buffer) Find(text []rune) (i int, j int) {
	for i, v := range b.data {
		j := strFind(v, text)
		if j >= 0 {
			return i, j
		}
	}
	return -1, -1
}
func (b *Buffer) Change(text, rp []rune, lineIdx int) {
	b.data[lineIdx] = (strReplace(b.data[lineIdx], text, rp))
}
