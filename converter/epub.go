package converter

import (
	"bookworm/scraper"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/bmaupin/go-epub"
)

func ToEpub(novel *scraper.Novel) (*epub.Epub, error) {

	// err := os.Remove((novel.Name + ".epub"))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	e := epub.NewEpub(novel.Name)

	e.SetAuthor(novel.Author)
	log.Println(novel.Name, novel.Author)

	coverCSSPath, _ := e.AddCSS("styles/css/cover.css", "")
	coverImagePath, _ := e.AddImage(novel.Cover, novel.Cover)
	e.SetCover(coverImagePath, coverCSSPath)

	_, err := e.AddFont("styles/fonts/redacted-script-regular.ttf", "font.ttf")
	if err != nil {
		log.Fatal(err)
	}

	var desc string
	for _, description := range novel.Description {
		desc += "<p>" + description + "</p>\n"
	}
	e.SetDescription(desc)

	for _, chapter := range novel.Chapters {

		sectionBody := "<h1>" + chapter.Title + "</h1>\n<p></p>\n"

		for _, paragraph := range chapter.Paragraph {
			sectionBody += "<p>" + paragraph + "</p>\n"
		}

		if _, err := e.AddSection(sectionBody, chapter.Title, "", ""); err != nil {
			return nil, err
		}
	}

	return e, nil
}

func safeClose(closer io.Closer) {
	if err := closer.Close(); err != nil {
		log.Panicf("error: %v", err)
	}
}

func ConvertToEpubFromJson(json_path string, epub_path string) {
	file, err := os.Open(json_path)
	if err != nil {
		log.Panicf("error: %v", err)
	}
	defer safeClose(file)

	byteResult, _ := ioutil.ReadAll(file)

	var novel scraper.Novel

	json.Unmarshal(byteResult, &novel)

	e, err := ToEpub(&novel)

	if err != nil {
		log.Panicf("error: %v", err)
	}

	log.Println("convert: done")

	log.Printf("writing %s to disk...", novel.Name+".epub")

	if err = e.Write(novel.Name + ".epub"); err != nil {
		log.Panicf("error: %v", err)
	}

	log.Println("write: done")
	log.Println("process: done")
}
