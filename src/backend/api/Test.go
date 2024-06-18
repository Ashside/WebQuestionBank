package api

type Test struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Author    string
	Questions []QuestionSummary `gorm:"foreignKey:TestId"`
}
