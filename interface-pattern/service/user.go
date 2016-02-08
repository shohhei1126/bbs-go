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

type UserImpl struct {
	userDao dao.User
}

func NewUser(userDao dao.User) User {
	return &UserImpl{userDao: userDao}
}

func (u UserImpl) Create(user *user.User) error {
	return u.userDao.Create(user)
}

func (u UserImpl) Update(user *user.User) error {
	return u.userDao.Update(user)
}

func (u UserImpl) FindById(id uint32) (user.User, error) {
	return u.userDao.FindById(id)
}

func (u UserImpl) FindByIds(ids []uint32) (user.UserSlice, error) {
	return u.userDao.FindByIds(ids)
}
