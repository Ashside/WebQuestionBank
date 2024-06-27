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
				text: "20世纪90年代，中国同东南亚各国关系进入全面发展时期。1997年，东南亚地区爆发金融危机，中国政府采取积极行动，对缓解危机、稳定东南亚地区经济乃至世界金融秩序作出了重要贡献。中共十八大以来，中国—东盟各方面关系全面加强：政治互信不断加强，经济合作硕果累累，人文交流日益频繁，中国“一带一路”倡议与东盟“互联互通总体规划发展”目标高度契合。中国—东盟“命运共同体”理念成为广泛共识。2021年11月，习近平指出：“中国东盟建立对话关系30年来……我们摆脱冷战阴霾，共同维护地区稳定。我们引领东亚经济一体化，促进共同发展繁荣，让20多亿民众过上了更好生活。”\n\n——摘编自黄庆、王巧荣主编《中华人民共和国外交史》等\n\n（1）根据材料一并结合所学知识，简析美国宣称放弃“门罗主义”的原因。（11分）\n\n（2）根据材料二并结合所学知识，概括20世纪90年代以来中国处理同东南亚国家关系的原则。（8分）\n\n（3）根据材料并结合所学知识，说明20世纪90年代以来中、美处理同周边国家关系的根本区别。（6分）",
			},
			want: []keywordResponse{
				{
					Keyword: "test",
					Score:   0,
				},
				{
					Keyword: "This",
					Score:   0,
				},
				{
					Keyword: "case",
					Score:   0,
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
