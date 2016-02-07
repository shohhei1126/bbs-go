package service

import (
	"github.com/shohhei1126/bbs-go/interface-pattern/dao"
	"github.com/shohhei1126/bbs-go/interface-pattern/model/user"
)

type User interface {
	Create(user *user.User) error
	Update(user *user.User) error
	FindById(id uint32) (user.User, error)
	FindByIds(ids []uint32) (user.UserSlice, error)
}

type UserServiceImpl struct {
	userDao dao.UserDao
}

func (u UserServiceImpl) Create(user *user.User) error {
	return u.userDao.Create(user)
}

func (u UserServiceImpl) Update(user *user.User) error {
	return u.userDao.Update(user)
}

func (u UserServiceImpl) FindById(id uint32) (user.User, error) {
	return u.userDao.FindById(id)
}

func (u UserServiceImpl) FindByIds(ids []uint32) (user.UserSlice, error) {
	return u.userDao.FindByIds(ids)
}
