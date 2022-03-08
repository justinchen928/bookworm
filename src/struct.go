/*
Copyright © 2022 rfaychen justin928501@gmail.com

*/
package src

import (
	"fmt"
	"os"
)

type Chapter struct {
	Title     string
	paragraph []string
}

type Novel struct {
	Name        string
	Author      string
	Chapters    []Chapter
	Description []string
	Cover       string
}

func (novel Novel) toTxt(dest string) {
	file_path := fmt.Sprintf("%s/%s.txt", dest, novel.Name)
	f, err := os.OpenFile(file_path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	var text = ""
	for _, chapter := range novel.Chapters {
		text += chapter.Title
		text += "\n\n"
		for _, paragraph := range chapter.paragraph {
			text += paragraph
		}
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}
