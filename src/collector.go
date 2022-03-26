/*
Copyright © 2022 rfaychen justin928501@gmail.com

*/
package src

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
