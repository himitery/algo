package programmers

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
			name:      "Problem 389481",
			problemId: "389481",
			want: []model.Problem{
				{
					Question: []interface{}{
						30,
						[]interface{}{"d", "e", "bb", "aa", "ae"},
					},
					Answer: "ah",
				},
				{
					Question: []interface{}{
						7388,
						[]interface{}{"gqk", "kdn", "jxj", "jxi", "fug", "jxg", "ewq", "len", "bhc"},
					},
					Answer: "jxk",
				},
			},
			wantErr: false,
		},
		{
			name:      "Problem 388354",
			problemId: "388354",
			want: []model.Problem{
				{
					Question: []interface{}{
						[]interface{}{11, 9, 3, 2, 4, 6},
						[]interface{}{
							[]interface{}{9, 11},
							[]interface{}{2, 3},
							[]interface{}{6, 3},
							[]interface{}{3, 4},
						},
					},
					Answer: []interface{}{1, 0},
				},
				{
					Question: []interface{}{
						[]interface{}{9, 15, 14, 7, 6, 1, 2, 4, 5, 11, 8, 10},
						[]interface{}{
							[]interface{}{5, 14},
							[]interface{}{1, 4},
							[]interface{}{9, 11},
							[]interface{}{2, 15},
							[]interface{}{2, 5},
							[]interface{}{9, 7},
							[]interface{}{8, 1},
							[]interface{}{6, 4},
						},
					},
					Answer: []interface{}{2, 1},
				},
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

			for i := range got {
				if i >= len(tt.want) {
					t.Errorf("extra result at index %d: %+v", i, got[i])
					continue
				}

				gotQ := got[i].Question
				wantQ := tt.want[i].Question
				if !reflect.DeepEqual(gotQ, wantQ) {
					t.Errorf("index %d - Question mismatch:\n  got:  %#v\n  want: %#v", i, gotQ, wantQ)
				}

				gotA := got[i].Answer
				wantA := tt.want[i].Answer
				if !reflect.DeepEqual(gotA, wantA) {
					t.Errorf("index %d - Answer mismatch:\n  got:  %#v\n  want: %#v", i, gotA, wantA)
				}
			}

			//if tt.want == nil && len(got) != 0 {
			//	t.Errorf("GetById() for not found problemId got = %v, want empty slice or nil", got)
			//}
			//if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetById() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
