package dao

import (
	"github.com/shohhei1126/bbs-go/common/time"
	"github.com/shohhei1126/bbs-go/interface/model"
	"testing"
	"github.com/stretchr/testify/assert"
	"database/sql"
)

func TestCommentSaveAndDelete(t *testing.T) {
	dbMap.TruncateTables()
	now := time.Now()

	comment := model.Comment{
		UserId:    1,
		ThreadId:  1,
		Body:      "comment_body",
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := commentDao.Create(&comment); err != nil {
		t.Fatal(err)
	}

	actualComment, err := commentDao.FindById(comment.Id)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, comment, *actualComment, "")

	if err := commentDao.Delete(actualComment); err != nil {
		t.Fatal(err)
	}

	_, err = commentDao.FindById(comment.Id)
	assert.Equal(t, sql.ErrNoRows, err, "")
}
