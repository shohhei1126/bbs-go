package dao

import (
	"github.com/Masterminds/squirrel"
	"github.com/shohhei1126/bbs-go/interface/model"
	"gopkg.in/gorp.v1"
	"math"
)

type Thread interface {
	Create(thread *model.Thread) error
	Update(user *model.Thread) error
	FindById(id uint32) (*model.Thread, error)
	Increment(threadId uint32, count int) error
}

type ThreadImpl struct {
	dbm *gorp.DbMap
	dbs *gorp.DbMap
}

func (t ThreadImpl) Create(thread *model.Thread) error {
	return t.dbm.Insert(thread)
}

func (t ThreadImpl) Update(user *model.Thread) error {
	_, err := t.dbm.Update(user)
	return err
}

func (t ThreadImpl) FindById(id uint32) (*model.Thread, error) {
	sql, args, err := squirrel.Select("*").
		From("threads").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return nil, err
	}
	var thread model.Thread
	if err := t.dbs.SelectOne(&thread, sql, args...); err != nil {
		return nil, err
	}
	return &thread, nil
}

func (t ThreadImpl) Increment(threadId uint32, count int) error {
	if count == 0 {
		return nil
	}

	var value interface{}
	if count > 0 {
		value = squirrel.Expr("comment_count + ?", count)
	} else {
		subtractValue := uint(math.Abs(float64(count)))
		value = squirrel.Expr("IF(comment_count < ?, 0, comment_count - ?)", subtractValue, subtractValue)
	}
	sql, args, err := squirrel.Update("threads").
		Set("comment_count", value).
		Where("id = ?", threadId).
		ToSql()

	if err != nil {
		return err
	}
	_, err = t.dbm.Exec(sql, args...)
	return nil
}
