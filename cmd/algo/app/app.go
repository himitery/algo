package app

import (
	"algo/internal/adapter"
	"algo/internal/domain/service"
	"algo/pkg/logger"
	"go.uber.org/fx"
)

func New(invoke fx.Option) *fx.App {
	return fx.New(
		fx.NopLogger,
		fx.Invoke(logger.New),

		adapter.Module(),
		service.Module(),

		invoke,
	)
}
