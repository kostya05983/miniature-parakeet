package parse

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

type ParserLink struct {
	result string
}

const catalogAddress = ""

//Get Link address of specified action
func (parse ParserLink) GetAddressOfAction(name string) string {
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

	c.OnHTML("div[class=catalog__search__body]", func(element *colly.HTMLElement) {
		element.DOM.Nodes[0].FirstChild.NextSibling.Attr[5].Val = name
		element.DOM.Nodes[0].FirstChild.NextSibling.Attr[1].Val = name
		element.DOM.Nodes[0].FirstChild.NextSibling.Data = name

		cl := colly.NewCollector()

		cl.OnHTML("span[class=catalog__line]", func(e *colly.HTMLElement) {
			a := e.DOM.Nodes[0].FirstChild.NextSibling.Attr[0].Val //link
			title := e.DOM.Nodes[0].FirstChild.NextSibling.LastChild.PrevSibling.Attr[1].Val
			//fmt.Printf("%s,%s", title, name)
			if title == name {
				parse.result = a
				fmt.Println("Result", a)
			}
		})

		err:= cl.Visit(catalogAddress)
		if err != nil {
			print(err)
		}
	})

	c.OnHTML("span[class=catalog__line]", func(e *colly.HTMLElement) {
		a := e.DOM.Nodes[0].FirstChild.NextSibling.Attr[0].Val //link
		title := e.DOM.Nodes[0].FirstChild.NextSibling.LastChild.PrevSibling.Attr[1].Val
		//fmt.Printf("%s,%s", title, name)
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
