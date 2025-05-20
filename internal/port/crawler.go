package port

import "algo/internal/domain/model"

type Crawler interface {
	GetById(id string) ([]model.Problem, error)
}
