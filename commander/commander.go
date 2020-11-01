package command

import (
	"fmt"
	"log"

	bf "github.com/shizuku/moon_editor/buffer"
)

type Commander struct {
	buffer bf.Buffer
}

func New() Commander {
	return Commander{}
}
func (c *Commander) Loop() {
	for {
		fmt.Print("command>")
		var co string
		var _, err = fmt.Scan(&co)
		if err != nil {
			log.Fatal(err)
			continue
		}
		switch co {
		case "r":
			c.read()
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
		case "h":
			c.help()
			break
		case "q":
			return
		default:
			fmt.Println("unknow command.")
			break
		}
	}
}

func (c *Commander) read() {
	var fileName string
	fmt.Print("filename>")
	var _, err = fmt.Scan(&fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.buffer.Read(fileName)
	c.buffer.Print()
}
func (c *Commander) write() {
	err := c.buffer.Write()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("writen.")
}
func (c *Commander) insert() {
}
func (c *Commander) delete() {
}
func (c *Commander) find() {
}
func (c *Commander) help() {
	fmt.Println("help:")
}
