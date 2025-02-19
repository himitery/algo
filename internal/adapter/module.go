package adapter

import (
	"algo/internal/adapter/baekjoon"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(baekjoon.New),
	)
}
