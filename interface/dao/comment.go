package dao

import (
	"github.com/Masterminds/squirrel"
	"github.com/shohhei1126/bbs-go/interface/model"
	"gopkg.in/gorp.v1"
)

type Comment interface {
	Create(thread *model.Comment) error
	Delete(user *model.Comment) error
	FindById(id uint32) (*model.Comment, error)
}

type CommentImpl struct {
	dbm *gorp.DbMap
	dbs *gorp.DbMap
}

func (c CommentImpl) Create(comment *model.Comment) error {
	return c.dbm.Insert(comment)
}

func (c CommentImpl) Delete(comment *model.Comment) error {
	_, err := c.dbm.Delete(comment)
	return err
}

func (c CommentImpl) FindById(id uint32) (*model.Comment, error) {
	sql, args, err := squirrel.Select("*").
		From("comments").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return nil, err
	}
	var comment model.Comment
	if err := c.dbs.SelectOne(&comment, sql, args...); err != nil {
		return nil, err
	}
	return &comment, nil
}