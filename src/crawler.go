/*
Copyright © 2022 rfaychen justin928501@gmail.com

*/
package src

func Crawler(first_page_link string) Novel {
	novel := Novel{}
	novel.Chapters = make([]Chapter, 0)
	shuchengCrawler(first_page_link, &novel)
	return novel
}

func SaveToFile(novel Novel, dest string) {
	novel.toTxt(dest)
}
