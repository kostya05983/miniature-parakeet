package parse

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

type Parser struct {
	result string
}

const catalogAddress = "https://quote.rbc.ru/catalog/"

func (parse Parser) GetAddressOfAction(name string) string {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", catalogAddress)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("span[class=catalog__line]", func(e *colly.HTMLElement) {
		a := e.DOM.Nodes[0].FirstChild.NextSibling.Attr[0].Val //link
		title := e.DOM.Nodes[0].FirstChild.NextSibling.LastChild.PrevSibling.Attr[1].Val
		if title == name {
			parse.result = a
			fmt.Println("Result", a)
		}
	})

	c.OnHTML("div[class=l-table js-filter-lock-block]", func(e *colly.HTMLElement) {
		fmt.Println("Found g-relative")
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished")
	})

	err := c.Visit(catalogAddress)
	if err != nil {
		print(err)
	}

	return parse.result
}
