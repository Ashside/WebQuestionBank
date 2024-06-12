package api

type Paper struct {
	Id                 int `gorm:"primaryKey"`
	ChoiceQuestion     []ChoiceQuestions
	SubjectiveQuestion []SubjectiveQuestions
}
