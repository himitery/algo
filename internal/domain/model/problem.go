package model

type Problem struct {
	Question interface{} `json:"inputs"`
	Answer   interface{} `json:"expected"`
}
