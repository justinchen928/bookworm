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

func (crawler ShuChengCrawler) Crawl(novel Novel) {

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
		r, _ := regexp.Compile(`(http|https):\/\/.*.html$`)
		log.Println("Link", chapter.Title, crawler.link, r.MatchString(crawler.link))
		if r.MatchString(crawler.link) {
			novel.Chapters = append(novel.Chapters, chapter)
			if len(novel.Chapters) >= 10 {
				novel.AppendToTxt()
				novel.Chapters = nil
				novel.Chapters = make([]Chapter, 0)
			}
			element.Request.Visit(crawler.link)
		} else {
			novel.Chapters = append(novel.Chapters, chapter)
			novel.AppendToTxt()
			log.Println("end")
		}
	})
	crawler.collector.Visit(crawler.link)
	crawler.collector.Wait()
}
