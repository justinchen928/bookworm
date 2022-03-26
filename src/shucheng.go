/*
Copyright © 2022 rfaychen justin928501@gmail.com

*/
package src

import (
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

type ShuChengCrawler struct {
	link      string
	collector *colly.Collector
}

func (crawler *ShuChengCrawler) SetCollector(collector *colly.Collector) {
	crawler.collector = collector
}

func isLastPage(link string) bool {
	r, _ := regexp.Compile(`(http|https):\/\/.*.html$`)
	return !r.MatchString(link)
}

func (crawler ShuChengCrawler) Crawl(novel *Novel) {

	chapter := novel.NewChapter()

	crawler.collector.OnHTML(".info > a", func(element *colly.HTMLElement) {
		novel.Name = element.Text
	})

	crawler.collector.OnHTML(".content > h1", func(element *colly.HTMLElement) {
		chapter.Title = element.Text
	})

	crawler.collector.OnHTML(".content > div.neirong", func(element *colly.HTMLElement) {
		element.ForEach("p", func(_ int, p_element *colly.HTMLElement) {
			paragraph := strings.TrimSpace(p_element.Text)
			reg := regexp.MustCompile(`(<|{|<| )`)
			paragraph = reg.ReplaceAllString(paragraph, "")
			chapter.paragraph = append(chapter.paragraph, paragraph)
			chapter.paragraph = append(chapter.paragraph, "\n\n")
		})
	})

	crawler.collector.OnHTML("#BookNext", func(element *colly.HTMLElement) {
		crawler.link = strings.TrimSpace(element.Attr("href"))
		novel.Chapters = append(novel.Chapters, chapter)
		log.Println(chapter.Title, crawler.link)
		chapter = novel.NewChapter()
		if isLastPage(crawler.link) {
			log.Println("end")
		} else {
			element.Request.Visit(crawler.link)
		}
	})
	crawler.collector.Visit(crawler.link)
	crawler.collector.Wait()
}
