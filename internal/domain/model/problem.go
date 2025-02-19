package model

type Problem struct {
	Question []string `json:"input"`
	Answer   []string `json:"expected"`
}
