package dao

import (
	"database/sql"
	"fmt"
	"github.com/shohhei1126/bbs-go/common/bbstime"
	"github.com/shohhei1126/bbs-go/interface/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCommentSaveAndDelete(t *testing.T) {
	dbMap.TruncateTables()
	now := bbstime.Now()

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

func TestCommentFindSliceByThreadId(t *testing.T) {
	dbMap.TruncateTables()

	// preparing data
	now := bbstime.Now()
	comments := make(model.CommentSlice, 10)
	for i := range comments {
		now = now.Add(time.Hour)
		comments[i].UserId = uint32(i + 1)
		comments[i].ThreadId = 1
		comments[i].Body = fmt.Sprintf("body %v", i)
		comments[i].CreatedAt = now
		comments[i].UpdatedAt = now
		if err := commentDao.Create(&comments[i]); err != nil {
			t.Fatal(err)
		}
	}

	tests := []struct {
		paging   Paging
		expected model.CommentSlice
	}{
		{
			paging: Paging{Limit: 3, Offset: 0, OrderBy: "created_at desc"},
			expected: comments[7:].SortBy(func(c1, c2 model.Comment) bool {
				return c1.CreatedAt.After(c2.CreatedAt)
			}),
		},
		{
			paging: Paging{Limit: 3, Offset: 3, OrderBy: "created_at desc"},
			expected: comments[4:7].SortBy(func(c1, c2 model.Comment) bool {
				return c1.CreatedAt.After(c2.CreatedAt)
			}),
		},
		{
			paging:   Paging{Limit: 3, Offset: 0, OrderBy: "id asc"},
			expected: comments[0:3],
		},
		{
			paging:   Paging{Limit: 3, Offset: 6, OrderBy: "id asc"},
			expected: comments[6:9],
		},
	}

	for _, test := range tests {
		actualComments, err := commentDao.FindSliceByThreadId(1, test.paging)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, test.expected, actualComments, "")
	}
}
