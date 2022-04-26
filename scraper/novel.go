package scraper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Chapter struct {
	Title     string   `json:"Title"`
	Paragraph []string `json:"Paragraph"`
}

type Novel struct {
	Name        string    `json:"Name"`
	Author      string    `json:"Author"`
	Description []string  `json:"Description"`
	Cover       string    `json:"Cover"`
	Chapters    []Chapter `json:"Chapters"`
}

func NewNovel() *Novel {
	return &Novel{}
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
		for _, paragraph := range chapter.Paragraph {
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

func (novel Novel) SaveToJson(path string) {
	file, _ := json.MarshalIndent(novel, "", " ")
	_ = ioutil.WriteFile(path, file, 0600)
}
