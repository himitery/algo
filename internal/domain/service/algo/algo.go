package algo

import (
	"algo/internal/domain/model"
	"algo/internal/port"
	"algo/internal/usecase"
	"algo/pkg/file"
	"algo/pkg/logger"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type state struct {
	baekjoon port.Baekjoon
}

func New(baekjoon port.Baekjoon) usecase.Algo {
	return &state{
		baekjoon: baekjoon,
	}
}

func (cls *state) Add(lang, platform, id string) {
	logger.Info(
		"Fetching problem data",
		zap.String("lang", lang), zap.String("platform", platform), zap.String("id", id),
	)

	problem := lo.Switch[string, []model.Problem](platform).
		CaseF("baekjoon", func() []model.Problem {
			return cls.baekjoon.GetById(id)
		}).
		DefaultF(func() []model.Problem {
			return nil
		})

	if problem == nil {
		logger.Error("No problem data found", zap.String("platform", platform), zap.String("id", id))
		return
	}

	logger.Info("Problem data retrieved successfully", zap.Any("problem", problem))

	data, err := json.MarshalIndent(problem, "", "  ")
	if err != nil {
		logger.Error("Error serializing problem data", zap.Error(err))
		return
	}

	path := fmt.Sprintf("%s/%s/%s", platform, lang, id)
	err = file.Save(path, "sample.json", data)
	if err != nil {
		logger.Error("Error saving problem file", zap.String("path", path+"/sample.json"), zap.Error(err))
		return
	}

	logger.Info("Problem file saved successfully", zap.String("path", path+"/sample.json"))
}
