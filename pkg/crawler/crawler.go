package crawler

import (
	"algo/internal/domain/model"
	"fmt"
	"github.com/gocolly/colly/v2"
)

const Agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"

func Execute(url, query string, callback func(*colly.HTMLElement) (interface{}, interface{})) ([]model.Problem, error) {
	var result []model.Problem
	var questions, answers []interface{}

	collector := colly.NewCollector(colly.UserAgent(Agent))

	collector.OnHTML(query, func(element *colly.HTMLElement) {
		question, answer := callback(element)
		if question != nil {
			questions = append(questions, question)
		}
		if answer != nil {
			answers = append(answers, answer)
		}
	})

	if err := collector.Visit(url); err != nil {
		return result, fmt.Errorf("failed to visit url (%s) : %w", url, err)
	}

	for idx := 0; idx < len(questions) && idx < len(answers); idx++ {
		result = append(result, model.Problem{
			Question: questions[idx],
			Answer:   answers[idx],
		})
	}

	return result, nil
}
