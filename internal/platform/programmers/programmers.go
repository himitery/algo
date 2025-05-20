package programmers

import (
	"algo/internal/domain/model"
	"algo/internal/port"
	"algo/pkg/crawler"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/samber/lo"
	"math"
	"strings"
)

type state struct {
	baseUrl   string
	baseQuery string
}

func New() port.Crawler {
	return &state{
		"https://school.programmers.co.kr/learn/courses/30/lessons/",
		"h5:contains('입출력 예') + table tbody tr",
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

	var values []string
	element.ForEach("td", func(_ int, col *colly.HTMLElement) {
		values = append(values, col.Text)
	})

	question, err := parse("[" + strings.Join(values[0:len(values)-1], ",") + "]")
	if err != nil {
		println("failed")
	}

	answer, err = parse(values[len(values)-1])
	if err != nil {
		println("failed 2, ", values[len(values)-1])
	}

	return question, answer
}

func parse(text string) (interface{}, error) {
	var parsed interface{}
	if err := json.Unmarshal([]byte(text), &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse answer: %w", err)
	}

	return normalize(parsed).(interface{}), nil
}

func normalize(v interface{}) interface{} {
	switch x := v.(type) {
	case float64:
		if x == math.Trunc(x) {
			return int(x)
		}

		return x

	case []interface{}:
		return lo.Map(x, func(elem interface{}, _ int) interface{} {
			return normalize(elem)
		})

	case map[string]interface{}:
		return lo.MapValues(x, func(val interface{}, _ string) interface{} {
			return normalize(val)
		})

	default:
		return x
	}
}
