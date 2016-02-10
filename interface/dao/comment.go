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
	FindSliceByThreadId(threadId uint32, paging Paging) (model.CommentSlice, error)
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

func (c CommentImpl) FindSliceByThreadId(threadId uint32, paging Paging) (model.CommentSlice, error) {
	builder := squirrel.Select("*").
		From("comments").
		Where("thread_id = ?", threadId).
		Limit(paging.Limit).
		Offset(paging.Offset)
	if paging.OrderBy != "" {
		builder = builder.OrderBy(paging.OrderBy)
	}
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	var comments model.CommentSlice
	if _, err := c.dbs.Select(&comments, sql, args...); err != nil {
		return nil, err
	}
	return comments, nil
}
