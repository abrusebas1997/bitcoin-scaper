package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	type coin struct {
		name   string
		price  int
		change string
	}

	var bitcoin coin

	// On every a element which has href attribute call callback
	c.OnHTML("#__layout > div > div.layout__wrp > div.header-zone.layout__header > header > div > div.tickers-desktop.header-desktop__tickers > ul > li:nth-child(1) > a > span.tickers-desktop__coin-cap", func(e *colly.HTMLElement) {
		bitcoin.name = e.Text
	})

	c.OnHTML("#__layout > div > div.layout__wrp > div.header-zone.layout__header > header > div > div.tickers-desktop.header-desktop__tickers > ul > li:nth-child(1) > a > span.tickers-desktop__coin-value", func(e *colly.HTMLElement) {

		// remove comma and $ sign
		str_num := strings.Replace(e.Text[2:], ",", "", -1)
		// trim white space and convert to int
		p, _ := strconv.Atoi(str_num[:len(str_num)-1])

		bitcoin.price = p
	})

	c.OnHTML("#__layout > div > div.layout__wrp > div.header-zone.layout__header > header > div > div.tickers-desktop.header-desktop__tickers > ul > li:nth-child(1) > a > span.tickers-desktop__coin-diff.tickers-desktop__coin-value_down", func(e *colly.HTMLElement) {
		bitcoin.change = e.Text
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://cointelegraph.com/tags/bitcoin")

	fmt.Println(bitcoin)

}
