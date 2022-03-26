package src

import (
	"fmt"
	"log"
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

func (novel Novel) NewChapter() Chapter {
	return Chapter{}
}

func (novel Novel) SaveToTxt(path string) {
	file_path := fmt.Sprintf("%s/%s.txt", path, novel.Name)
	f, err := os.OpenFile(file_path, os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var text = ""
	var start = ""
	var end = ""
	for index, chapter := range novel.Chapters {
		if start == "" {
			start = chapter.Title
		}
		end = chapter.Title
		text += chapter.Title
		text += "\n\n"
		for _, paragraph := range chapter.paragraph {
			text += paragraph
		}

		if index%20 == 0 {
			log.Println("write from", start, "to", end)
			if _, err = f.WriteString(text); err != nil {
				panic(err)
			}
			start = ""
			text = ""
		}
	}

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}
