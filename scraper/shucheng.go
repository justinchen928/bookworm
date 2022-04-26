/*
Copyright © 2022 rfaychen justin928501@gmail.com

*/
package scraper

import (
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

type ShuChengScraper struct {
	novel     *Novel
	URL       string
	collector *colly.Collector
}

func isLastPage(link string) bool {
	r, _ := regexp.Compile(`(http|https):\/\/.*.html$`)
	return !r.MatchString(link)
}

func (ss *ShuChengScraper) GetNovel() *Novel {
	return ss.novel
}

func (ss *ShuChengScraper) Start() {

	chapter := ss.novel.NewChapter()

	ss.collector.OnHTML(".info > a", func(element *colly.HTMLElement) {
		ss.novel.Name = element.Text
	})

	ss.collector.OnHTML(".content > h1", func(element *colly.HTMLElement) {
		chapter.Title = element.Text
	})

	ss.collector.OnHTML(".content > div.neirong", func(element *colly.HTMLElement) {
		element.ForEach("p", func(_ int, p_element *colly.HTMLElement) {
			paragraph := strings.TrimSpace(p_element.Text)
			reg := regexp.MustCompile(`(<|{|<| )`)
			paragraph = reg.ReplaceAllString(paragraph, "")
			chapter.Paragraph = append(chapter.Paragraph, paragraph)
		})
	})

	ss.collector.OnHTML("#BookNext", func(element *colly.HTMLElement) {
		ss.URL = strings.TrimSpace(element.Attr("href"))
		ss.novel.Chapters = append(ss.novel.Chapters, chapter)
		log.Println(chapter.Title, ss.URL)
		chapter = ss.novel.NewChapter()
		if isLastPage(ss.URL) {
			log.Println("end")
		} else {
			element.Request.Visit(ss.URL)
		}
	})
	ss.collector.Visit(ss.URL)
	ss.collector.Wait()
}
