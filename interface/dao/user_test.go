package dao

import (
	"github.com/shohhei1126/bbs-go/common/time"
	"github.com/shohhei1126/bbs-go/interface/model"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserSaveUpdate(t *testing.T) {
	dbMap.TruncateTables()
	now := time.Now()

	user := model.User{
		Username:    "username1",
		Password:    "password1",
		DisplayName: "displayName1",
		Status:      model.Member,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := userDao.Create(&user); err != nil {
		t.Fatal(err)
	}

	user.DisplayName = "displayname1_updated"
	user.UpdatedAt = now
	if err := userDao.Update(&user); err != nil {
		t.Fatal(err)
	}

	actualUser, err := userDao.FindById(user.Id)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, user, *actualUser, "")
}

func TestUserFindById(t *testing.T) {
	dbMap.TruncateTables()
	now := time.Now()

	u1 := model.User{
		Username:    "username1",
		Password:    "password1",
		DisplayName: "displayName1",
		Status:      model.Member,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := userDao.Create(&u1); err != nil {
		t.Fatal(err)
	}

	actualUser, err := userDao.FindById(u1.Id)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, u1, *actualUser, "")
}

func TestUserFindByIds(t *testing.T) {
	dbMap.TruncateTables()
	now := time.Now()

	u1 := model.User{
		Username:    "username1",
		Password:    "password1",
		DisplayName: "displayName1",
		Status:      model.Member,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	u2 := model.User{
		Username:    "username2",
		Password:    "password2",
		DisplayName: "displayName2",
		Status:      model.Member,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := userDao.Create(&u1); err != nil {
		t.Fatal(err)
	}
	if err := userDao.Create(&u2); err != nil {
		t.Fatal(err)
	}

	actualUsers, err := userDao.FindByIds([]uint32{u1.Id, u2.Id})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(actualUsers), "")
	assert.Equal(t, u1, actualUsers[0], "")
	assert.Equal(t, u2, actualUsers[1], "")
}
