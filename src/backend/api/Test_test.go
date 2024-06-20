package api

import (
	"gorm.io/gorm"
	"log"
	"reflect"
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
			got, err := GenerateMdByTestID(tt.args.db, tt.args.id)
			log.Println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateMdByTestID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateMdByTestID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryTestByID(t *testing.T) {
	type args struct {
		db *gorm.DB
		id int
	}
	dbT, _ := getDatabase()
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{db: dbT, id: 2},
			want: []int{6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := queryTestByID(tt.args.db, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryTestByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
