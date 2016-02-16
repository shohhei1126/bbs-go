package dao

import (
	"github.com/Masterminds/squirrel"
	"github.com/shohhei1126/bbs-go/model"
	"gopkg.in/gorp.v1"
)

type Thread interface {
	Create(thread *model.Thread) error
	FindList(paging Paging) (model.ThreadSlice, error)
}

type ThreadImpl struct {
	dbm *gorp.DbMap
	dbs *gorp.DbMap
}

func NewThread(dbm, dbs *gorp.DbMap) Thread {
	return &ThreadImpl{dbm: dbm, dbs: dbs}
}

func (t ThreadImpl) Create(thread *model.Thread) error {
	return t.dbm.Insert(thread)
}

func (t ThreadImpl) FindList(paging Paging) (model.ThreadSlice, error) {
	sql, args, err := squirrel.
		Select("*").
		From("threads").
		OrderBy(paging.OrderBy).
		Limit(paging.Limit).
		Offset(paging.Offset).
		ToSql()
	if err != nil {
		return nil, err
	}
	var threads model.ThreadSlice
	if _, err := t.dbs.Select(&threads, sql, args...); err != nil {
		return nil, err
	}
	return threads, nil
}
