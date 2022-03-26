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
	SavePath    string
}

func (novel Novel) NewChapter() Chapter {
	return Chapter{}
}

func (novel Novel) AppendToTxt() {
	file_path := fmt.Sprintf("%s/%s.txt", novel.SavePath, novel.Name)
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
