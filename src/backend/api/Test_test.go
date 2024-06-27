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
			if got := QueryQuesIdByTestID(tt.args.db, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryQuesIdByTestID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteTestByID(t *testing.T) {
	type args struct {
		db *gorm.DB
		i  int
	}
	dbT, _ := getDatabase()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "1",
			args:    args{db: dbT, i: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteTestByID(tt.args.db, tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTestByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
