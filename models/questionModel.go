package models

import (
	"github.com/Kamva/mgm/v3"
)

// Question is the model that defines a question entry
type Question struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string   `json:"title" bson:"title"`
	InputType        string   `json:"inputType" bson:"inputType"`
	InputSpec        string   `json:"inputSpec" bson:"inputSpec"`
	Choices          []Choice `json:"choices" bson:"choices"`
}

// Choice is the model that defines a choice entry
type Choice struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
}

/*
// CreateQuestion is a wrapper that creates a new question entry
func CreateQuestion(title, inputType, inputSpec string) *Question {
	return &Question{
		Title:		title,
		InputType:	inputType,
		InputSpec:	inputSpec,
		//todo: Choices
	}
}
*/
