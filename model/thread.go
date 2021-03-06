package model

import (
	"time"
)

// +gen slice:"Select[uint32],SortBy"
type Thread struct {
	Id        uint32    `json:"id" db:"id"`
	UserId    uint32    `json:"-" db:"user_id"`
	Title     string    `json:"title" db:"title"`
	Body      string    `json:"body" db:"body"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`

	User *User `json:"user" db:"-"`
}
