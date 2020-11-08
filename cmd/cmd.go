package cmd

import (
	"fmt"
	"log"

	bf "github.com/shizuku/moon_editor/buffer"
	st "github.com/shizuku/moon_editor/string"
)

type Commander struct {
	opened   bool
	fileName string
	buffer   bf.Buffer
	cursor   int
}

func New() Commander {
	return Commander{
		opened:   false,
		fileName: "",
		buffer:   bf.New(),
		cursor:   1,
	}
}
func (c *Commander) Loop() {
	for {
		for !c.opened {
			c.Read()
		}
		fmt.Printf("%s; %dL >", c.fileName, c.cursor)
		var co string
		var _, err = fmt.Scan(&co)
		if err != nil {
			log.Fatal(err)
			continue
		}
		switch co {
		case "r":
			c.Read()
			break
		case "w":
			c.write()
			break
		case "i":
			c.insert()
			break
		case "d":
			c.delete()
			break
		case "f":
			c.find()
			break
		case "c":
			c.change()
			break
		case "q":
			return
		case "h":
			c.help()
			break
		case "n":
			c.next()
			break
		case "p":
			c.prev()
			break
		case "b":
			c.begin()
			break
		case "e":
			c.end()
			break
		case "g":
			c.gotoo()
			break
		case "v":
			c.view()
			break
		default:
			fmt.Println("unknow command.")
			break
		}
	}
}

func (c *Commander) Read() {
	var fileName string
	fmt.Print("filename>")
	var _, err = fmt.Scan(&fileName)
	if err != nil {
		log.Println(err)
		return
	}
	err = c.buffer.Read(fileName)
	if err != nil {
		log.Println(err)
		return
	}
	c.opened = true
	c.fileName = fileName
	c.view()
}
func (c *Commander) write() {
	err := c.buffer.Write(c.fileName)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("writen.")
}
func (c *Commander) insert() {
	var (
		text string
	)
	fmt.Print("text to insert>")
	_, err := fmt.Scan(&text)
	if err != nil {
		log.Println(err)
		return
	}
	c.buffer.Insert(st.New([]rune(text)), c.cursor-1)
}
func (c *Commander) delete() {
	c.cursor = c.buffer.Delete(c.cursor-1) + 1
}
func (c *Commander) find() {
	var (
		text string
	)
	fmt.Print("text to find>")
	_, err := fmt.Scan(&text)
	if err != nil {
		log.Fatalln(err)
		return
	}
	i, j := c.buffer.Find(st.New([]rune(text)), c.cursor-1)
	fmt.Printf("%dL, %dC\n", i+1, j+1)

}
func (c *Commander) change() {
	var (
		text string
		rp   string
	)
	fmt.Print("text to find>")
	_, err := fmt.Scan(&text)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Print("text to replace>")
	_, err = fmt.Scan(&rp)
	if err != nil {
		log.Println(err)
		return
	}
	c.buffer.Change(st.New([]rune(text)), st.New([]rune(rp)), c.cursor-1)
}
func (c *Commander) help() {
	fmt.Println("help:")
}
func (c *Commander) next() {
	c.cursor++
	if c.cursor > c.buffer.LineNum() {
		c.end()
	}
}
func (c *Commander) prev() {
	c.cursor--
	if c.cursor < 1 {
		c.begin()
	}
}
func (c *Commander) begin() {
	c.cursor = 1
}
func (c *Commander) end() {
	c.cursor = c.buffer.LineNum()
}
func (c *Commander) gotoo() {
	fmt.Print("line index>")
	var lineIdx int
	_, err := fmt.Scan(&lineIdx)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c.cursor = lineIdx
	if c.cursor < 0 {
		c.begin()
	}
	if c.cursor >= c.buffer.LineNum() {
		c.end()
	}
}
func (c *Commander) view() {
	c.buffer.Print()
}
