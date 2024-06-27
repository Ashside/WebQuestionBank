package api

import (
	"gorm.io/gorm"
	"testing"
)

func TestCheckScore(t *testing.T) {
	type args struct {
		db     *gorm.DB
		assign Assignments
	}
	dbT, _ := getDatabase()
	asT, _ := GetAssignsByStuName(dbT, "zck")
	tests := []struct {
		name string
		args args
		want float64
	}{

		{
			name: "Test1",
			args: args{
				db:     dbT,
				assign: asT[0],
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckScore(tt.args.db, tt.args.assign); got != tt.want {
				t.Errorf("CheckScore() = %v, want %v", got, tt.want)
			}
		})
	}
}