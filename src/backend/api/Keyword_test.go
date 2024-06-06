package api

import (
	"reflect"
	"testing"
)

func Test_getKeyword(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    []keywordResponse
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{text: "学习书法，就选唐颜真卿《颜勤礼碑》原碑与对临「第1节」"},
			want: []keywordResponse{
				{keyword: "颜勤礼碑", score: 0.2669559478242008},
				{keyword: "书法", score: 0.25579790927326845},
				{keyword: "原碑", score: 0.23924662247807055},
				{keyword: "唐颜真卿", score: 0.23799952042446024},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getKeyword(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("getKeyword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKeyword() got = %v, want %v", got, tt.want)
			}
		})
	}
}
