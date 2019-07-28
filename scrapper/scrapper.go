package scrapper

import (
	"fmt"

	"github.com/gocolly/colly"
)

func InitializeScrapper() *colly.Collector {
	c := colly.NewCollector()

	// Find and visit all links
	// c.OnHTML("a", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// 	fmt.Printf("%s", e.Response.Body)
	// })

	c.OnHTML("div.WarframeNavBoxImage2 > a", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))
		// e.ChildAttr()
		fmt.Printf("%s", e.Response.Body)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(e *colly.Response) {
		fmt.Printf("%s", e.Body)
	})

	c.Visit("https://warframe.fandom.com/wiki/Warframes")

	return c
}

//to check if the content type is image and to save it
// if strings.Index(r.Headers.Get("Content-Type"), "image") > -1 {
// 	r.Save(outputDir + r.FileName())
// 	return
// }
