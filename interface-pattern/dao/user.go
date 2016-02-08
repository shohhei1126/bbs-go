package dao

import (
	"database/sql"
	sq "github.com/lann/squirrel"
	"github.com/shohhei1126/bbs-go/interface-pattern/model/user"
	"gopkg.in/gorp.v1"
)

type User interface {
	Create(user *user.User) error
	Update(user *user.User) error
	FindById(id uint32) (user.User, error)
	FindByIds(ids []uint32) (user.UserSlice, error)
}

type UserImpl struct {
	dbm *gorp.DbMap
	dbs *gorp.DbMap
}

func NewUser(dbm, dbs *gorp.DbMap) User {
	return &UserImpl{dbm: dbm, dbs: dbs}
}

func (u UserImpl) Create(user *user.User) error {
	return u.dbm.Insert(user)
}

func (u UserImpl) Update(user *user.User) error {
	_, err := u.dbm.Update(user)
	return err
}

func (u UserImpl) FindById(id uint32) (user.User, error) {
	users, err := u.FindByIds([]uint32{id})
	if err != nil {
		return user.User{}, err
	}
	if len(users) == 0 {
		return user.User{}, sql.ErrNoRows
	}
	return users[0], nil
}

func (u UserImpl) FindByIds(ids []uint32) (user.UserSlice, error) {
	if len(ids) == 0 {
		return user.UserSlice{}, nil
	}
	sql, args, err := sq.Select("*").
		From("users").
		Where(sq.Eq{"id": ids}).
		OrderBy("id ASC").
		ToSql()
	if err != nil {
		return nil, err
	}
	var users user.UserSlice
	_, err = u.dbs.Select(&users, sql, args...)
	if err != nil {
		return nil, err
	}
	return users, nil
}
