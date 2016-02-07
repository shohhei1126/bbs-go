package user

import "time"

// +gen slice:"Select[uint32]"
type User struct {
	Id          uint32    `db:"id"`
	Username    string    `db:"username"`
	Password    string    `db:"password"`
	DisplayName string    `db:"display_name"`
	Status      Status    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type Status uint8

const (
	Member Status = iota + 1
	Withdrawal
)
