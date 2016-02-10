package model

import "time"

// +gen slice:"Select[uint32]"
type Comment struct {
	Id        uint32    `json:"id" db:"id"`
	UserId    uint32    `json:"userId" db:"user_id"`
	ThreadId  uint32    `json:"threadId" db:"thread_id"`
	Body      string    `json:"body" db:"body"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}