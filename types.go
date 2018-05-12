package main

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Card struct {
	CardID          int
	Title           string
	CardType        string
	DescriptionBox1 string
	DescriptionBox2 string
	DescriptionBox3 string
	Priority        int
	Severity        int
	AssignedTo      string
	CaseNumber      int
	Column          int
}

type Comment struct {
	CommentID int
	CardID    int
	Author    string
	Message   string
	Posted    string
}

func (c Comment) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.CardID, validation.Required),
		validation.Field(&c.Author, validation.Required),
		validation.Field(&c.Message, validation.Required, validation.Length(1, 3000)))
}

func (c Card) Validate() error {
	// Validate fields for Card struct
	return validation.ValidateStruct(&c,
		// title under 140 chars
		validation.Field(&c.Title, validation.Required, validation.Length(2, 140)),
		// correct type
		validation.Field(&c.CardType, validation.Required, validation.In("Story", "Bug", "Feature",
			"Epic").Error("accepted values:\"Story\", \"Bug\", \"Feature\", \"Epic\"")),
		validation.Field(&c.Priority, validation.Required, validation.In(1, 2, 3, 4,
			5).Error("Priority must be between 1 and 5")),
		validation.Field(&c.Severity, validation.Required, validation.In(1, 2,
			3).Error("Severity must be between 1 and 3")),
		// UTFletternumeric fails if there are spaces
		validation.Field(&c.AssignedTo, validation.Required, is.UTFLetterNumeric),
		//
		validation.Field(&c.Column, validation.Required, validation.In(1, 2, 3, 4, 5, 6, 7, 8, 9,
			10, 11, 12).Error("Invalid Column ID")))
}
