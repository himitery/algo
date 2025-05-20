package usecase

import (
	"algo/internal/platform/baekjoon"
	"algo/internal/platform/programmers"
	"algo/internal/port"
	"algo/pkg/file"
	"fmt"
	"github.com/samber/lo"
)

func Add(platform, id, path string) error {
	crawler := lo.Switch[string, port.Crawler](platform).
		Case("baekjoon", baekjoon.New()).
		Case("programmers", programmers.New()).
		Default(nil)
	if crawler == nil {
		return fmt.Errorf("platform not supported: %s", platform) // 수정: 오류 메시지에 플랫폼 이름 포함
	}

	problems, err := crawler.GetById(id)
	if err != nil {
		return fmt.Errorf("failed to get problems: %w", err)
	}

	return file.Save(path, problems)
}
