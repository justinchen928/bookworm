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

func shuchengCrawler(first_page_link string, novel *Novel) {
	domain := "www.51shucheng.net"
	chapter := Chapter{}

	collector := colly.NewCollector(
		colly.Async(true),
		colly.AllowedDomains(domain),
	)

	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
	})

	collector.OnRequest(func(request *colly.Request) {
		chapter = Chapter{}
	})

	collector.OnHTML(".info > a", func(element *colly.HTMLElement) {
		novel.Name = element.Text
	})

	collector.OnHTML(".content > h1", func(element *colly.HTMLElement) {
		chapter.Title = element.Text
	})

	collector.OnHTML(".content > div.neirong", func(element *colly.HTMLElement) {
		element.ForEach("p", func(_ int, p_element *colly.HTMLElement) {
			paragraph := strings.TrimSpace(p_element.Text)
			reg := regexp.MustCompile(`(<|{|<| )`)
			paragraph = reg.ReplaceAllString(paragraph, "")
			chapter.paragraph = append(chapter.paragraph, paragraph)
			chapter.paragraph = append(chapter.paragraph, "\n\n")
		})
	})

	collector.OnHTML("#BookNext", func(element *colly.HTMLElement) {
		next_page_link := strings.TrimSpace(element.Attr("href"))
		r, _ := regexp.Compile(`(http|https):\/\/.*.html$`)
		log.Println("Link", chapter.Title, next_page_link, r.MatchString(next_page_link))
		if r.MatchString(next_page_link) {
			novel.Chapters = append(novel.Chapters, chapter)
			if len(novel.Chapters) >= 10 {
				// novel.toTxt(os.Args[1])
				novel.Chapters = nil
				novel.Chapters = make([]Chapter, 0)
			}
			element.Request.Visit(next_page_link)
		} else {
			novel.Chapters = append(novel.Chapters, chapter)
			log.Println("end")
		}
	})
	collector.Visit(first_page_link)
	collector.Wait()
}
