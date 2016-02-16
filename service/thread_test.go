package service

import (
	"github.com/golang/mock/gomock"
	"github.com/shohhei1126/bbs-go/dao"
	"github.com/shohhei1126/bbs-go/model"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestThreadFindThreads(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	paging := dao.Paging{OrderBy: "updated_at", Limit: 3, Offset: 3}
	threads := model.ThreadSlice{
		{Id: 2, UserId: 12},
		{Id: 3, UserId: 13},
		{Id: 4, UserId: 14},
	}
	threadDaoMoc := dao.NewMockThread(ctl)
	threadDaoMoc.EXPECT().FindList(paging).Return(threads, nil)

	users := model.UserSlice{
		{Id: 12, Username: "username 12"},
		{Id: 13, Username: "username 13"},
		{Id: 14, Username: "username 14"},
	}
	userDaoMoc := dao.NewMockUser(ctl)
	userDaoMoc.EXPECT().FindByIds([]uint32{12, 13, 14}).Return(users, nil)

	threadService := NewThread(userDaoMoc, threadDaoMoc)

	actualThreads, err := threadService.FindThreads(paging)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int(paging.Limit), len(actualThreads), "")

	for _, thread := range actualThreads {
		assert.Equal(t, thread.UserId, thread.User.Id)
	}
}