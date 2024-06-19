package api

import (
	"gorm.io/gorm"
	"testing"
)

func Test_findAvailableTestsId(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	dbT, _ := getDatabase()
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "Test1",
			args: args{db: dbT},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAvailableTestsId(tt.args.db); got != tt.want {
				t.Errorf("findAvailableTestsId() = %v, want %v", got, tt.want)
			}
		})
	}
}
