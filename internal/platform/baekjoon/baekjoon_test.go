package baekjoon

import (
	"algo/internal/domain/model"
	"reflect"
	"testing"
)

func TestGetById(t *testing.T) {
	tests := []struct {
		name      string
		problemId string
		want      []model.Problem
		wantErr   bool
	}{
		{
			name:      "Problem 1000",
			problemId: "1000",
			want: []model.Problem{
				{Question: []string{"1 2"}, Answer: []string{"3"}},
			},
			wantErr: false,
		},
		{
			name:      "Problem 10000",
			problemId: "10000",
			want: []model.Problem{
				{Question: []string{"2", "1 3", "5 1"}, Answer: []string{"3"}},
				{Question: []string{"3", "2 2", "1 1", "3 1"}, Answer: []string{"5"}},
				{Question: []string{"4", "7 5", "-9 11", "11 9", "0 20"}, Answer: []string{"6"}},
			},
			wantErr: false,
		},
		{
			name:      "Problem 99999 (not found)",
			problemId: "99999",
			want:      nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := New()
			got, err := app.GetById(tt.problemId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want == nil && len(got) != 0 {
				t.Errorf("GetById() for not found problemId got = %v, want empty slice or nil", got)
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
