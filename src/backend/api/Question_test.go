package api

import (
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestQueryQuestionFromId(t *testing.T) {
	type args struct {
		db *gorm.DB
		id int
	}
	dbTest, _ := getDatabase()
	tests := []struct {
		name    string
		args    args
		want    ChoiceQuestions
		wantErr bool
	}{
		// TODO: Add test cases.
		{

			name: "Test1",
			args: args{
				db: dbTest,
				id: 1,
			},
			want: ChoiceQuestions{
				Id:         1,
				Subject:    "history",
				Content:    "阅读下面的材料，根据要求写作。（60分）\n\n\t随着互联网的普及、人工智能的应用，越来越多的问题能很快得到答案。那么，我们的问题是否会越来越少？\n\n\t以上材料引发了你怎样的联想和思考？请写一篇文章。\n\n\t要求：选准角度，确定立意，明确文体，自拟标题；不要套作，不得抄袭；不得泄露个人信息；不少于800字。",
				Answer:     "略",
				Options:    "",
				Difficulty: "1",
				Author:     "admin@hit.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryQuestionFromId(tt.args.db, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryQuestionFromId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryQuestionFromId() got = %v, want %v", got, tt.want)
			}
		})
	}
}
