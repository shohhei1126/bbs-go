package model

import (
	"errors"
	"time"
)

// +gen slice:"Select[uint32]"
type User struct {
	Id          uint32     `json:"id" db:"id"`
	Username    string     `json:"username" db:"username"`
	Password    string     `json:"password" db:"password"`
	DisplayName string     `json:"displayName" db:"display_name"`
	Status      UserStatus `json:"status" db:"status"`
	CreatedAt   time.Time  `json:"createAt" db:"created_at"`
	UpdatedAt   time.Time  `json:"updatedAt" db:"updated_at"`
}

type UserStatus uint8

const (
	Member UserStatus = iota + 1
	Withdrawal
)

func (s UserStatus) MarshalJSON() ([]byte, error) {
	status := ""
	switch s {
	case Member:
		status = "member"
	case Withdrawal:
		status = "withdrawal"
	default:
		return nil, errors.New("unsuported user status parameter")
	}
	return []byte(`"` + status + `"`), nil
}
