package api

import (
	"reflect"
	"testing"
)

func Test_getKeywordFromLocal(t *testing.T) {
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
			args: args{
				text: "This is a test case",
			},
			want: []keywordResponse{
				{
					Keyword: "test case",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getKeywordFromLocal(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("getKeywordFromLocal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKeywordFromLocal() got = %v, want %v", got, tt.want)
			}
		})
	}
}
