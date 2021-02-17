package models

import (
	"github.com/Kamva/mgm/v3"
)

// Survey is the model that defines a survey entry
type Survey struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"title" bson:"title"`
	Token            string `json:"token" bson:"token"`
	IsDeleted        bool   `json:"isDeleted" bson:"isDeleted"`
	Questions		 []Question `json:"questions" bson:"questions"`
}

// CreateSurvey is a wrapper that creates a new survey entry
func CreateSurvey(title, token string, questions []Question) *Survey {
	return &Survey{
		Title:     title,
		Token:     token,
		IsDeleted: false,
		Questions: questions,
	}
}
