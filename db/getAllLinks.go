package db

import (
	"github.com/pkg/errors"
	"link-logger/db/models"
	"link-logger/interfaces"
)

// GetAllLinks returns all links
func GetAllLinks() ([]interfaces.Page, error) {
	result := make([]*models.Links, 0)
	err := Get().Find(&result).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	pages := make([]interfaces.Page, len(result))
	for i := range result {
		pages[i] = result[i]
	}

	return pages, nil
}
