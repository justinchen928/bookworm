package src

import (
	"regexp"

	"github.com/gocolly/colly"
)

type Crawler interface {
	SetCollector(collector *colly.Collector)
	Crawl(novel *Novel)
}

func NewBookCrawler(link string) Crawler {
	if regexp.MustCompile(`www.51shucheng.net`).MatchString(link) {
		return &ShuChengCrawler{link: link}
	}
	return nil
}
