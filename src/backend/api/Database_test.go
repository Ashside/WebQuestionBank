package api

import (
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func Test_getDatabase(t *testing.T) {
	tests := []struct {
		name    string
		want    *gorm.DB
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test getDatabase",
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getDatabase()
			if (err != nil) != tt.wantErr {
				t.Errorf("getDatabase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDatabase() got = %v, want %v", got, tt.want)
			}
		})
	}
}
