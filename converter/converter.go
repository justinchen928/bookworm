package converter

import (
	"bookworm/scraper"
	"fmt"
	"log"
	"os"
)

type Config struct {
	OutputDir string
}

type Converter struct {
	config Config
}

func New(config Config) *Converter {
	return &Converter{config: config}
}

func (c *Converter) ConvertAndSave(novel *scraper.Novel, format string) {

	if "epub" == format {
		location := fmt.Sprintf("%s%s.%s", c.config.OutputDir, novel.Name, format)
		e, err := ToEpub(novel)
		if err != nil {
			log.Panicf("error: %v", err)
		}
		if err = e.Write(location); err != nil {
			log.Panicf("error: %v", err)
		}
		log.Println("convert: done")
		log.Printf("writing %s to %s", novel.Name, location)
	}

	if "kepub" == format {
		epub_location := fmt.Sprintf("%s%s.%s", c.config.OutputDir, novel.Name, "epub")
		e, err := ToEpub(novel)
		if err != nil {
			log.Panicf("error: %v", err)
		}
		if err = e.Write(epub_location); err != nil {
			log.Panicf("error: %v", err)
		}

		kepub_location := fmt.Sprintf("%s%s.%s.epub", c.config.OutputDir, novel.Name, format)
		ToKepub(epub_location, kepub_location)
		log.Println("convert: done")
		log.Printf("writing %s to %s", novel.Name, kepub_location)

		err = os.Remove(epub_location)
		if err != nil {
			log.Fatal(err)
		}
	}

	if "json" == format {
		location := fmt.Sprintf("%s%s.%s", c.config.OutputDir, novel.Name, format)
		novel.SaveToJson(location)
		log.Println("convert: done")
		log.Printf("writing %s to %s", novel.Name, location)
	}
}
