package api

import (
	"gorm.io/gorm"
	"testing"
)

func Test_getQuestionTypeByID(t *testing.T) {
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
		// TODO: Add test cases.
		{
			name:    "Test1",
			args:    args{db: dbT, id: 1},
			want:    "subjective",
			wantErr: false,
		},
		{
			name:    "Test2",
			args:    args{db: dbT, id: 6},
			want:    "choice",
			wantErr: false,
		},
		{
			name:    "Test3",
			args:    args{db: dbT, id: 14},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getQuestionTypeByID(tt.args.db, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getQuestionTypeByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getQuestionTypeByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
