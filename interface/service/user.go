package service

import (
	"github.com/shohhei1126/bbs-go/interface/dao"
	"github.com/shohhei1126/bbs-go/interface/model"
)

type User interface {
	Create(user *model.User) error
	Update(user *model.User) error
	FindById(id uint32) (*model.User, error)
	FindByIds(ids []uint32) (model.UserSlice, error)
}

type UserImpl struct {
	userDao dao.User
}

func NewUser(userDao dao.User) User {
	return &UserImpl{userDao: userDao}
}

func (u UserImpl) Create(user *model.User) error {
	return u.userDao.Create(user)
}

func (u UserImpl) Update(user *model.User) error {
	return u.userDao.Update(user)
}

func (u UserImpl) FindById(id uint32) (*model.User, error) {
	return u.userDao.FindById(id)
}

func (u UserImpl) FindByIds(ids []uint32) (model.UserSlice, error) {
	return u.userDao.FindByIds(ids)
}
