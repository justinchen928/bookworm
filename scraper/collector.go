/*
Copyright © 2024 Justin Chen justin928501@gmail.com

*/
package scraper

import (
	"github.com/gocolly/colly"
)

func NewCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.Async(true),
	)
	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
	})
	return collector
}
