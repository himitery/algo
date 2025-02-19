package service

import (
	"algo/internal/domain/service/algo"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(algo.New),
	)
}
