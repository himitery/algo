package baekjoon

import (
	"algo/internal/domain/model"
	"algo/internal/port"
	"algo/pkg/crawler"
	"algo/pkg/logger"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"strings"
)

type state struct {
	crawler *colly.Collector
	url     string
}

func New() port.Baekjoon {
	return &state{
		crawler: crawler.New(),
		url:     "https://www.acmicpc.net/problem/",
	}
}

func (cls *state) GetById(id string) []model.Problem {
	var problems []model.Problem
	var question, answer []string

	cls.crawler.OnHTML("pre.sampledata", func(element *colly.HTMLElement) {
		id := element.Attr("id")
		if strings.Contains(id, "input") {
			question = append(question, element.Text)
		}
		if strings.Contains(id, "output") {
			answer = append(answer, element.Text)
		}
	})

	if err := cls.crawler.Visit(cls.url + id); err != nil {
		logger.Error("Failed to visit baekjoon: %v", zap.Error(err))

		return nil
	}

	for idx := range question {
		problems = append(problems, model.Problem{
			Question: strings.Split(strings.TrimSpace(question[idx]), "\n"),
			Answer:   strings.Split(strings.TrimSpace(answer[idx]), "\n"),
		})
	}

	return problems
}
