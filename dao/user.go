package dao

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/shohhei1126/bbs-go/model"
	"gopkg.in/gorp.v1"
)

type User interface {
	Create(user *model.User) error
	FindById(id uint32) (*model.User, error)
	FindByIds(ids []uint32) (model.UserSlice, error)
}

type UserImpl struct {
	dbm *gorp.DbMap
	dbs *gorp.DbMap
}

func NewUser(dbm, dbs *gorp.DbMap) User {
	return &UserImpl{dbm: dbm, dbs: dbs}
}

func (u UserImpl) Create(user *model.User) error {
	return u.dbm.Insert(user)
}

func (u UserImpl) FindById(id uint32) (*model.User, error) {
	users, err := u.FindByIds([]uint32{id})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, sql.ErrNoRows
	}
	return &users[0], nil
}

func (u UserImpl) FindByIds(ids []uint32) (model.UserSlice, error) {
	if len(ids) == 0 {
		return model.UserSlice{}, nil
	}
	sql, args, err := squirrel.Select("*").
		From("users").
		Where(squirrel.Eq{"id": ids}).
		OrderBy("id ASC").
		ToSql()
	if err != nil {
		return nil, err
	}
	var users model.UserSlice
	_, err = u.dbs.Select(&users, sql, args...)
	if err != nil {
		return nil, err
	}
	return users, nil
}
