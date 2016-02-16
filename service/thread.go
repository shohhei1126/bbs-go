package service

import (
	"github.com/shohhei1126/bbs-go/dao"
	"github.com/shohhei1126/bbs-go/model"
)

type Thread interface {
	FindThreads(paging dao.Paging) (model.ThreadSlice, error)
}

type ThreadImpl struct {
	userDao    dao.User
	threadDao  dao.Thread
	commentDao dao.Comment
}

func NewThread(userDao dao.User, thread dao.Thread, commentDao dao.Comment) Thread {
	return ThreadImpl{userDao: userDao, threadDao: thread, commentDao: commentDao}
}

func (t ThreadImpl) FindThreads(paging dao.Paging) (model.ThreadSlice, error) {
	threads, err := t.threadDao.FindList(paging)
	if err != nil {
		return nil, err
	}
	userIds := threads.SelectUint32(func(t model.Thread) uint32 {
		return t.UserId
	})

	users, err := t.userDao.FindByIds(userIds)
	if err != nil {
		return nil, err
	}

	userMap := users.GroupByUint32(func(u model.User) uint32 {
		return u.Id
	})

	for i := range threads {
		user, ok := userMap[threads[i].UserId]
		if ok == false || len(user) == 0 {
			continue
		}
		threads[i].User = &user[0]
	}
	return threads, nil
}
