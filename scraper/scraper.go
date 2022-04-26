package scraper

import (
	"regexp"
)

type Config struct {
	URL string
}

type Scraper interface {
	Start()
	GetNovel() *Novel
}

func New(config Config) Scraper {
	novel := NewNovel()
	collector := NewCollector()
	if regexp.MustCompile(`www.51shucheng.net`).MatchString(config.URL) {
		return &ShuChengScraper{
			novel:     novel,
			URL:       config.URL,
			collector: collector,
		}
	}
	return nil
}
