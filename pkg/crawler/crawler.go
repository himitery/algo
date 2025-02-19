package crawler

import "github.com/gocolly/colly/v2"

var agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"

func New() *colly.Collector {
	return colly.NewCollector(
		colly.UserAgent(agent),
	)
}
