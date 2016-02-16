package dao

import (
	"github.com/shohhei1126/bbs-go/bbstime"
	"github.com/shohhei1126/bbs-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestThreadCreate(t *testing.T) {
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
	actualThread := model.Thread{}
	dbMap.SelectOne(&actualThread, "select * from threads where id = ?", thread.Id)
	assert.Equal(t, thread, actualThread, "")
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
