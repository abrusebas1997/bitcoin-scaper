package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type coin struct {
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Change string `json:"change"`
	Description string `json:"description"`
}

func writeFile(file []byte) {
	this := ioutil.WriteFile("result.json", file, 0644)
	if err := this; err != nil {
		panic(err)
	}
}

func serializeJSON(foo coin) {
	fmt.Println("Serializing Data", foo)
	btcJson, _ := json.Marshal(foo)
	writeFile(btcJson)
	fmt.Print("Serializing Complete ")
	fmt.Println(string(btcJson))
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	var bitcoin coin

	// On every a element which has href attribute call callback
	c.OnHTML("#__layout > div > div.layout__wrp > div.header-zone.layout__header > header > div > div.tickers-desktop.header-desktop__tickers > ul > li:nth-child(1) > a > span.tickers-desktop__coin-cap", func(e *colly.HTMLElement) {
		bitcoin.Name = e.Text
	})

	c.OnHTML("#__layout > div > div.layout__wrp > div.header-zone.layout__header > header > div > div.tickers-desktop.header-desktop__tickers > ul > li:nth-child(1) > a > span.tickers-desktop__coin-value", func(e *colly.HTMLElement) {

		// remove comma and $ sign
		str_num := strings.Replace(e.Text[2:], ",", "", -1)
		// trim white space and convert to int
		p, _ := strconv.Atoi(str_num[:len(str_num)-1])

		bitcoin.Price = p
	})

	c.OnHTML("#__layout > div > div.layout__wrp > div.header-zone.layout__header > header > div > div.tickers-desktop.header-desktop__tickers > ul > li:nth-child(1) > a > span.tickers-desktop__coin-diff.tickers-desktop__coin-value_down", func(e *colly.HTMLElement) {
		bitcoin.Change = e.Text
	})

	c.OnHTML("#__layout > div > div.layout__wrp > main > div > div > div.tag-about.tag-page__about > div.tag-about__desc-col > div > p", func(e *colly.HTMLElement) {
		bitcoin.Description = e.Text
	})
	

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://cointelegraph.com/tags/bitcoin")

	serializeJSON(bitcoin)

}
