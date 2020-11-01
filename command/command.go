package command

import (
	"fmt"
	"log"

	"github.com/shizuku/moon_editor/editor"
)

func read(e editor.Editor) {
	var fileName string
	fmt.Print("filename>")
	var _, err = fmt.Scan(&fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	e.Read(fileName)
	e.Print()
}

func Loop() {
	var e = editor.New()
	for {
		fmt.Print("command>")
		var c string
		var _, err = fmt.Scan(&c)
		if err != nil {
			log.Fatal(err)
			continue
		}
		switch c {
		case "r":
			read(e)
		case "w":
		}
	}
}
