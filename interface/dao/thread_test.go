package dao

import (
	"github.com/shohhei1126/bbs-go/common/time"
	"github.com/shohhei1126/bbs-go/interface/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreadSaveUpdate(t *testing.T) {
	dbMap.TruncateTables()
	now := time.Now()

	thread := model.Thread{
		UserId:    1,
		Title:     "thread_title",
		Body:      "thread_body",
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := threadDao.Create(&thread); err != nil {
		t.Fatal(err)
	}

	thread.Title = "thread_title_updated"
	thread.Body = "thread_body_updated"
	if err := threadDao.Update(&thread); err != nil {
		t.Fatal(err)
	}

	threadActual, err := threadDao.FindById(thread.Id)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, thread, *threadActual, "")
}

func TestThreadIncrementCommentCount(t *testing.T) {
	dbMap.TruncateTables()

	testData := []struct {
		first     uint32
		increment int
		expected  uint32
	}{
		{first: 0, increment: 1, expected: 1},
		{first: 5, increment: -1, expected: 4},
		{first: 3, increment: -5, expected: 0},
		{first: 3, increment: 7, expected: 10},
	}

	for _, data := range testData {
		thread := model.Thread{CommentCount: data.first}
		if err := threadDao.Create(&thread); err != nil {
			t.Fatal(err)
		}
		if err := threadDao.Increment(thread.Id, data.increment); err != nil {
			t.Fatal(err)
		}
		threadActual, err := threadDao.FindById(thread.Id)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, data.expected, threadActual.CommentCount, "")
	}

}