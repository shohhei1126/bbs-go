package dao

import (
	"github.com/shohhei1126/bbs-go/common/time"
	"github.com/shohhei1126/bbs-go/interface-pattern/model/user"

	"reflect"
	"testing"
)

func TestUserUpdate(t *testing.T) {
	dbMap.TruncateTables()

	userDao := UserDaoImpl{dbm: dbMap, dbs: dbMap}
	now := time.Now()

	u1 := user.User{
		Username:    "username1",
		Password:    "password1",
		DisplayName: "displayName1",
		Status:      user.Member,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := userDao.Create(&u1); err != nil {
		t.Fatal("u1 create failed.", err)
	}

	u1.DisplayName = "displayname1_updated"
	u1.UpdatedAt = now
	if err := userDao.Update(&u1); err != nil {
		t.Fatal("u1 update failed.", err)
	}

	actualUser, err := userDao.FindById(u1.Id)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(actualUser, u1) == false {
		t.Errorf("actualUser and u1 are not equal. \n%+v\n%+v", actualUser, u1)
	}

}

func TestUserFindById(t *testing.T) {
	dbMap.TruncateTables()

	userDao := UserDaoImpl{dbm: dbMap, dbs: dbMap}
	now := time.Now()

	u1 := user.User{
		Username:    "username1",
		Password:    "password1",
		DisplayName: "displayName1",
		Status:      user.Member,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := userDao.Create(&u1); err != nil {
		t.Fatal("u1 create failed.", err)
	}

	actualUser, err := userDao.FindById(u1.Id)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(actualUser, u1) == false {
		t.Errorf("actualUser and u1 are not equal. \n%+v\n%+v", actualUser, u1)
	}
}

func TestUserFindByIds(t *testing.T) {
	dbMap.TruncateTables()

	userDao := UserDaoImpl{dbm: dbMap, dbs: dbMap}
	now := time.Now()

	u1 := user.User{
		Username:    "username1",
		Password:    "password1",
		DisplayName: "displayName1",
		Status:      user.Member,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	u2 := user.User{
		Username:    "username2",
		Password:    "password2",
		DisplayName: "displayName2",
		Status:      user.Member,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := userDao.Create(&u1); err != nil {
		t.Fatal("u1 create failed.", err)
	}
	if err := userDao.Create(&u2); err != nil {
		t.Fatal("u2 create failed.", err)
	}

	actualUsers, err := userDao.FindByIds([]uint32{u1.Id, u2.Id})
	if err != nil {
		t.Fatal(err)
	}

	if len(actualUsers) != 2 {
		t.Fatalf("len(actualUsers) = %v, want 2", len(actualUsers))
	}

	if reflect.DeepEqual(actualUsers[0], u1) == false {
		t.Errorf("actualUsers[0] and u1 are not equal. \n%+v\n%+v", actualUsers[0], u1)
	}

	if reflect.DeepEqual(actualUsers[1], u2) == false {
		t.Errorf("actualUsers[1] and u2 are not equal. \n%+v\n%+v", actualUsers[1], u2)
	}
}
