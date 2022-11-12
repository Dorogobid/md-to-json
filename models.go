package main

type Question struct {
	CategoryID		uint	`json:"category_id"`
	QuestionText	string	`json:"question_text"`
	CodeBlock		string	`json:"code_block"`
	CodeType		string	`json:"code_type"`
	Answers			[]Answer	`json:"answers"`
}

type Answer struct {
	AnswerText		string	`json:"answer_text"`
	IsTrue			bool	`json:"is_true"`
}