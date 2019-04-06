package parse

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strconv"
)

type AverageParser struct {
	result float64
	amount int
}

//Count average from all actions on page
func (parser AverageParser) GetAverage(address string, month string, dayPrev int) float64 {
	c := colly.NewCollector()

	prevMonth := parser.getMonthFromStr(month)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", address)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("span[class=item__review__container]", func(e *colly.HTMLElement) {
		r, _ := regexp.Compile("\\d+")
		data := e.DOM.Nodes[0].FirstChild.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling.FirstChild.Data
		var date string
		if e.DOM.Nodes[0].PrevSibling.PrevSibling != nil {
			date = e.DOM.Nodes[0].PrevSibling.PrevSibling.FirstChild.NextSibling.FirstChild.Data
		} else {
			date = e.DOM.Nodes[0].Parent.PrevSibling.PrevSibling.FirstChild.NextSibling.FirstChild.Data
		}

		day := date[:2]
		dayDigit, err := strconv.Atoi(day)
		if err != nil {
			fmt.Println(err)
		}

		month := date[3:9]
		monthDigit := parser.getMonthFromStr(month)

		if prevMonth > monthDigit || prevMonth == monthDigit && dayPrev > dayDigit {
			return
		}

		f, err := strconv.ParseFloat(r.FindString(data), 64)
		if err != nil {
			fmt.Println(err)
		}
		parser.result += f
		parser.amount += 1
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished")
	})

	err := c.Visit(address)
	if err != nil {
		print(err)
	}

	return parser.result / float64(parser.amount)
}

//Get month digit from str
func (parser AverageParser) getMonthFromStr(month string) int {
	switch month {
	case "янв":
		return 0
	case "фев":
		return 1
	case "мар":
		return 2
	case "апр":
		return 3
	case "май":
		return 4
	case "июн":
		return 5
	case "июл":
		return 6
	case "авг":
		return 7
	case "сен":
		return 8
	case "окт":
		return 9
	case "ноя":
		return 10
	case "дек":
		return 11
	default:
		return 0
	}
}
