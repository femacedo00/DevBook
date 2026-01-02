package models

import (
	"errors"
	"strings"
	"time"
)

type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedIn  time.Time `json:"createdIn,omitempty"`
}

// Prepare uses validate and format methods in the received publication
func (publication *Publication) Prepare() error {
	if error := publication.validation(); error != nil {
		return error
	}

	publication.format()
	return nil
}

func (publication *Publication) validation() error {
	if publication.Title == "" {
		return errors.New("Title is required and cannot be blank")
	}

	if publication.Content == "" {
		return errors.New("Content is required and cannot be blank")
	}

	return nil
}

func (publication *Publication) format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}
