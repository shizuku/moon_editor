package buffer

import (
	"bufio"
	"fmt"
	"io"
	"os"

	u "github.com/shizuku/moon_editor/unicode"
)

type Buffer struct {
	fileName string
	data     []u.String
}

func New() Buffer {
	return Buffer{}
}
func (e *Buffer) Read(fileName string) error {
	e.fileName = fileName
	var s u.String
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		return err
	}
	r := bufio.NewReader(f)
	for {
		ch, _, er := r.ReadRune()
		if er != nil {
			if er == io.EOF {
				break
			} else {
				return er
			}
		}
		s.Append(ch)
	}
	e.data = s.Split('\n')
	return nil
}
func (e *Buffer) String() string {
	var r []rune
	for _, v := range e.data {
		r = append(r, []rune(v.String())...)
		r = append(r, '\n')
	}
	return string(r)
}
func (e *Buffer) Print() {
	fmt.Println("------------------------------------------------------")
	fmt.Print(e.String())
	fmt.Println("------------------------------------------------------")
}
func (e *Buffer) Write() error {
	f, err := os.Open(e.fileName)
	defer f.Close()
	if err != nil {
		return err
	}
	f.WriteString(e.String())
	return nil
}
