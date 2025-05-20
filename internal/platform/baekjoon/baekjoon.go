package baekjoon

import (
	"algo/internal/domain/model"
	"algo/internal/port"
	"algo/pkg/crawler"
	"github.com/gocolly/colly/v2"
	"strings"
)

type state struct {
	baseUrl   string
	baseQuery string
}

func New() port.Crawler {
	return &state{
		"https://www.acmicpc.net/problem/",
		"pre.sampledata",
	}
}

func (c *state) GetById(id string) ([]model.Problem, error) {
	return crawler.Execute(
		c.baseUrl+id,
		c.baseQuery,
		callback,
	)
}

func callback(element *colly.HTMLElement) (interface{}, interface{}) {
	var question, answer interface{}

	id := element.Attr("id")
	text := strings.TrimSpace(element.Text)

	if strings.Contains(id, "input") {
		question = strings.Split(text, "\n")
	}
	if strings.Contains(id, "output") {
		answer = strings.Split(text, "\n")
	}

	return question, answer
}
