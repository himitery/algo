package port

import "algo/internal/domain/model"

type Baekjoon interface {
	GetById(id string) []model.Problem
}
