package dao

import (
	"github.com/shohhei1126/bbs-go/common/bbstime"
	"github.com/shohhei1126/bbs-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestThreadSaveUpdate(t *testing.T) {
	dbMap.TruncateTables()
	now := bbstime.Now()

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

func TestThreadFindList(t *testing.T) {
	dbMap.TruncateTables()
	createdAt := bbstime.Now()
	updatedAt := bbstime.Now()
	threads := make(model.ThreadSlice, 10)
	for i := range threads {
		createdAt = createdAt.Add(time.Hour)
		updatedAt = updatedAt.Add(-time.Hour)
		threads[i].CreatedAt = createdAt
		threads[i].UpdatedAt = updatedAt
		if err := dbMap.Insert(&threads[i]); err != nil {
			t.Fatal(err)
		}
	}

	tests := []struct {
		paging   Paging
		expected model.ThreadSlice
	}{
		{
			paging: Paging{OrderBy: "created_at desc", Limit: 3, Offset: 0},
			expected: threads[7:].SortBy(func(t1, t2 model.Thread) bool {
				return t1.CreatedAt.After(t2.CreatedAt)
			})},
		{
			paging: Paging{OrderBy: "created_at desc", Limit: 3, Offset: 6},
			expected: threads[1:4].SortBy(func(t1, t2 model.Thread) bool {
				return t1.CreatedAt.After(t2.CreatedAt)
			})},
		{
			paging:   Paging{OrderBy: "updated_at desc", Limit: 3, Offset: 0},
			expected: threads[0:3],
		},
		{
			paging:   Paging{OrderBy: "updated_at desc", Limit: 3, Offset: 6},
			expected: threads[6:9],
		},
	}

	for _, test := range tests {
		threads, err := threadDao.FindList(test.paging)
		t.Logf("%+v", threads)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, test.expected, threads, "")
	}
}
