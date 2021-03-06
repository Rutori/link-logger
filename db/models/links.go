package models

import (
	"gorm.io/gorm"
)

type Links struct {
	Title       string
	Address     string `gorm:"uniqueIndex"`
	Preview     string
	Description string

	gorm.Model
}

func (l Links) GetURL() string {
	return l.Address
}

func (l Links) GetTitle() string {
	return l.Title
}

func (l Links) GetPreview() string {
	return l.Preview
}

func (l Links) GetDescription() string {
	return l.Preview
}
