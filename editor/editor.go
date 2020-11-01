package editor

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/shizuku/moon_editor/str"
)

type Editor struct {
	fileName string
	data     []str.String
}

func New() Editor {
	return Editor{}
}

func (e *Editor) Read(fileName string) error {
	e.fileName = fileName
	var s str.String
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

func (e *Editor) String() string {
	var r []rune
	for _, v := range e.data {
		r = append(r, []rune(v.String())...)
		r = append(r, '\n')
	}
	return string(r)
}

func (e *Editor) Print() {
	fmt.Println(e.String())
}

func (e *Editor) Write() error {
	f, err := os.Open(e.fileName)
	defer f.Close()
	if err != nil {
		return err
	}
	f.WriteString(e.String())
	return nil
}
