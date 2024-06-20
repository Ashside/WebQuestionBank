package api

import (
	"gorm.io/gorm"
	"log"
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

func TestGeneratePDF(t *testing.T) {
	type args struct {
		db *gorm.DB
		id int
	}
	dbT, _ := getDatabase()
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Test1",
			args:    args{db: dbT, id: 1},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateMD(tt.args.db, tt.args.id)
			log.Println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateMD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateMD() got = %v, want %v", got, tt.want)
			}
		})
	}
}
